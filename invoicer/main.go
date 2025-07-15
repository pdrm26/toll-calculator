package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pdrm26/toll-calculator/types"
)

func writeJSON(w http.ResponseWriter, status int, res any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

func handleAggregate(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := service.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func handleInvoice(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obuID := r.URL.Query().Get("obu")
		fmt.Println(len(obuID))
		if len(obuID) == 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing obu id"})
			return
		}
		w.Write([]byte("return the invoice for an OBU"))
	}
}

func makeHTTPTransport(listenAddr string, service Aggregator) {
	fmt.Println("HTTP transport is running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(service))
	http.HandleFunc("/invoice", handleInvoice(service))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "http server listen address")
	flag.Parse()

	store := NewMemoryStore()
	service := NewInvoiceAggregator(store)
	service = NewLogMiddleware(service)

	makeHTTPTransport(*listenAddr, service)
}
