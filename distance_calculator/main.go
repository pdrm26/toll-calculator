package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pdrm26/toll-calculator/invoicer/client"
)

var (
	// httpListenAddr     = os.Getenv("AGG_HTTP_ENDPOINT")
	grpcListenAddr     = os.Getenv("AGG_GRPC_ENDPOINT")
	aggregatorEndpoint = fmt.Sprintf("localhost:%s", grpcListenAddr)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	service := NewCalculatorService()
	service = NewLogMiddleware(service)

	// httpClient := client.NewHTTPClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewkafkaConsumer(os.Getenv("KAFKA_TOPIC_NAME"), service, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
