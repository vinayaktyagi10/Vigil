package broker
import (
	"log"
	"encoding/json"
	queue "github.com/rabbitmq/amqp091-go"
	"vigil/internal/item"
	"context"
	"time"
)
const QueueName = "vigil.feed"

func Publish(channel *queue.Channel, feedItem item.Item){
	q, err := channel.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err == nil {
		log.Printf("Queue declared: %s", q.Name)
	} else {
		log.Printf("Error declaring queue: %s", err)
		return
	}

	body, err := json.Marshal(feedItem)
	if err != nil {
		log.Printf("Error marshalling feed item: %s", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		"",
		QueueName,
		false,
		false,
		queue.Publishing{
			ContentType: "application/json",
			Body:		 body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	}
}

func Consume(channel *queue.Channel) (<-chan queue.Delivery, error){
	q, err := channel.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err == nil {
		log.Printf("Queue declared: %s", q.Name)
	} else {
		log.Printf("Error declaring queue: %s", err)
		return nil, err
	}

	deliveries, err := channel.Consume(
		QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Error consuming queue: %s", err)
		return nil, err
	} else {
		return deliveries, nil
	}
}





