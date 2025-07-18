package main

import (
	"log"

	"github.com/pdrm26/toll-calculator/invoicer/client"
)

func main() {
	service := NewCalculatorService()
	service = NewLogMiddleware(service)

	// httpClient := client.NewHTTPClient("localhost:8001")
	grpcClient, err := client.NewGRPCClient("localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewkafkaConsumer("obudata", service, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
