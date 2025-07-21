package main

import (
	"log"

	"github.com/pdrm26/toll-calculator/invoicer/client"
)

const (
	aggregatorEndpoint = "localhost:4001"
)

func main() {
	service := NewCalculatorService()
	service = NewLogMiddleware(service)

	// httpClient := client.NewHTTPClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewkafkaConsumer("obudata", service, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
