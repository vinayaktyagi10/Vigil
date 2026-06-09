package main

import (
	"os"
	"vigil/internal/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	deliveries, err := broker.Consume(ch)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}

	for d := range deliveries {
		log.Printf("Recieved message: %s", d.Body)
	}
}
