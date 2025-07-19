package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/pdrm26/toll-calculator/invoicer/client"
	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "the listen address of the http server")
	aggregatorServiceAddr := flag.String("aggServiceAddr", "http://localhost:3000", "the listen address of the aggregator service")
	flag.Parse()
	client := client.NewHTTPClient(*aggregatorServiceAddr)
	invoice := newInvoiceHandler(client)

	http.HandleFunc("/invoice", makeAPIFunc(invoice.handleGetInvoice))
	logrus.Infof("gateway HTTP server is up and running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

type invoiceHandler struct {
	client client.Client
}

func newInvoiceHandler(client client.Client) *invoiceHandler {
	return &invoiceHandler{
		client: client,
	}
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri":  r.RequestURI,
			}).Info("REQ :: ")
		}(time.Now())
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func (i *invoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	invoice, err := i.client.GetInvoice(context.Background(), 8978773209462795273)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, invoice)
}

func writeJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
