package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pdrm26/toll-calculator/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) Aggregate(ctx context.Context, r *types.AggregatorDistance) error {
	distBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(distBytes))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error status %d", res.StatusCode)
	}

	return nil
}

func (c *HTTPClient) GetInvoice(ctx context.Context, obuID int) (*types.Invoice, error) {
	invoiceBody := &types.GetInvoiceRequets{
		Obuid: int64(obuID),
	}

	b, err := json.Marshal(invoiceBody)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s?obu=%d", c.Endpoint, "invoice", obuID)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status %d", res.StatusCode)
	}

	var inv types.Invoice
	if err := json.NewDecoder(res.Body).Decode(&inv); err != nil {
		return nil, err
	}

	return &inv, err

}
