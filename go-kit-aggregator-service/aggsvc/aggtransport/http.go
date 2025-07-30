package aggtransport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"

	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggendpoint"
	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggservice"
)

func NewHTTPHandler(endpoints aggendpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	m := http.NewServeMux()
	m.Handle("/aggregate", httptransport.NewServer(
		endpoints.AggregateEndpoint,
		decodeHTTPAggregateRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/invoice", httptransport.NewServer(
		endpoints.InvoiceEndpoint,
		decodeHTTPInvoiceRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return m
}

func NewHTTPClient(instance string, logger log.Logger) (aggservice.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	var options []httptransport.ClientOption
	var aggregateEndpoint endpoint.Endpoint
	{
		aggregateEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/aggregate"),
			encodeHTTPGenericRequest,
			decodeHTTPAggregateResponse,
			options...,
		).Endpoint()
	}

	var invoiceEndpoint endpoint.Endpoint
	{
		invoiceEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/invoice"),
			encodeHTTPGenericRequest,
			decodeHTTPInvoiceResponse,
			options...,
		).Endpoint()
	}

	return aggendpoint.Set{
		AggregateEndpoint: aggregateEndpoint,
		InvoiceEndpoint:   invoiceEndpoint,
	}, nil
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("ERROR ENCODER: ", err)
}

func decodeHTTPAggregateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req aggendpoint.AggregateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPInvoiceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req aggendpoint.InvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}

	r.Body = io.NopCloser(&buf)
	return nil
}

func decodeHTTPAggregateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp aggendpoint.AggregateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeHTTPInvoiceResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp aggendpoint.InvoiceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
