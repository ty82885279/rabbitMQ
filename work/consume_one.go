package main

import "rabbitMQ/pkg/rabbitMQ"

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQSimple("LLSimple")
	rabbitmq.ConsumeSimple()
}
