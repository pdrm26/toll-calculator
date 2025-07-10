package main

import (
	"log"
)

func main() {
	kafkaConsumer, err := NewkafkaConsumer()
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
