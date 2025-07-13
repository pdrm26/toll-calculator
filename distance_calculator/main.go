package main

import (
	"log"

	"github.com/pdrm26/toll-calculator/invoicer/client"
)

func main() {
	service := NewCalculatorService()
	service = NewLogMiddleware(service)
	client := client.NewClient("http://localhost:8000/aggregate")
	kafkaConsumer, err := NewkafkaConsumer("obudata", service, client)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
