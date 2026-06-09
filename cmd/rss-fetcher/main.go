package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"vigil/internal/broker"
	"vigil/internal/rss"
	"time"
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

	urls := []string{
		"https://news.ycombinator.com/rss",
	}
	for {

		for _, url := range urls {
			items, err := rss.Fetch(url)
			if err != nil {
				log.Printf("Failed to fetch RSS: %s", err)
				continue
			}

			for _, item := range items {
				broker.Publish(ch, item)
			}
		}
		time.Sleep(1 * time.Hour)
	}

}
