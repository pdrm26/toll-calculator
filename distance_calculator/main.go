package main

import (
	"log"
)

func main() {
	kafkaConsumer, err := NewkafkaConsumer("obudata")
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
