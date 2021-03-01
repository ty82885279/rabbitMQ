package main

import "rabbitMQ/pkg/rabbitMQ"

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQPubSub("exNewPub")
	rabbitmq.ConsumePub()
}
