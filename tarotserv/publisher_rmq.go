package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// RmqPublisher is the RabbitMQ publisher implementation.
type RmqPublisher struct {
	channel *amqp.Channel
	queue   string
}

// NewRmqPublisher creates new RmqPublisher instance.
func NewRmqPublisher(uri, queue string) (res *RmqPublisher, err error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return
	}

	channel, err := connection.Channel()

	return &RmqPublisher{
		channel: channel,
		queue:   queue,
	}, err
}

// Publish sends message to queue.
func (rp RmqPublisher) Publish(e Event) error {
	fmt.Println("Start publish event to rmq")

	bytes, err := json.Marshal(e)
	if err != nil {
		return err
	}

	return rp.channel.Publish(
		"tracking",
		"tracking",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
}

// CloseConnection close connection.
func (rp RmqPublisher) CloseConnection() {
	rp.channel.Close()
}
