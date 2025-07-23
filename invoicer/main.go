package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pdrm26/toll-calculator/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func writeJSON(w http.ResponseWriter, status int, res any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

func handleAggregate(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
			return
		}
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
		if r.Method != "GET" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
			return
		}
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

func makeGRPCTransport(listenAddr string, service Aggregator) error {
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer lis.Close()

	server := grpc.NewServer()
	types.RegisterAggregatorServer(server, NewGRPCServer(service))

	fmt.Println("✅ gRPC server running on", listenAddr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
	return nil
}

func makeHTTPTransport(listenAddr string, service Aggregator) {
	fmt.Println("HTTP transport is running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(service))
	http.HandleFunc("/invoice", handleInvoice(service))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func makeStore() Storer {
	storeType := os.Getenv("AGG_STORE_TYPE")
	switch storeType {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatalf("invalid memory type: %s", storeType)
		return nil
	}
}

func main() {
	httpListenAddr := os.Getenv("AGG_HTTP_ENDPOINT")
	grpcListenAddr := os.Getenv("AGG_GRPC_ENDPOINT")
	store := makeStore()
	service := NewInvoiceAggregator(store)
	service = NewMetricMiddleware(service)
	service = NewLogMiddleware(service)

	go makeGRPCTransport(grpcListenAddr, service)
	makeHTTPTransport(httpListenAddr, service)
}
