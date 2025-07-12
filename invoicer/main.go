package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pdrm26/toll-calculator/types"
)

func handleAggregate(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}

func makeHTTPTransport(listenAddr string, service Aggregator) {
	fmt.Println("HTTP transport is running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(service))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "http server listen address")
	flag.Parse()

	store := NewMemoryStore()
	service := NewInvoiceAggregator(store)

	makeHTTPTransport(*listenAddr, service)
}
