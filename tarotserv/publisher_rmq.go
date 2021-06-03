package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// RmqPublisher is the RabbitMQ publisher implementation.
type RmqPublisher struct {
	channel  *amqp.Channel
	exchange string
}

// NewRmqPublisher creates new RmqPublisher instance.
func NewRmqPublisher(uri, exchange string) (res *RmqPublisher, err error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return
	}

	channel, err := connection.Channel()

	return &RmqPublisher{
		channel:  channel,
		exchange: exchange,
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
		rp.exchange,
		rp.exchange,
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
