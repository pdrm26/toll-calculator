package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/pdrm26/toll-calculator/invoicer/client"
	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "the listen address of the http server")
	flag.Parse()

	client := client.NewHTTPClient("localhost:8000")
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
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func (i *invoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	invoice, err := i.client.GetInvoice(context.Background(), 10)
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
