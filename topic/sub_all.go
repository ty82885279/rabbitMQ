package main

import (
	"rabbitMQ/pkg/rabbitMQ"
)

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQTopic("exNewTopic", "#")
	rabbitmq.ConsumeTopic()
}
