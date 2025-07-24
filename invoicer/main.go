package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pdrm26/toll-calculator/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	aggMetricsHandler := NewHTTPMetricHandler("aggregate")
	invoiceMetricsHandler := NewHTTPMetricHandler("invoice")
	http.HandleFunc("/aggregate", aggMetricsHandler.instrument(handleAggregate(service)))
	http.HandleFunc("/invoice", invoiceMetricsHandler.instrument(handleInvoice(service)))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
