package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		obuParam := r.URL.Query().Get("obu")
		if len(obuParam) == 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing obu id"})
			return
		}

		obuID, err := strconv.Atoi(obuParam)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid obu id"})
			return
		}

		invoice, err := service.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, invoice)
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
