package main

import "rabbitMQ/pkg/rabbitMQ"

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQTopic("exNewTopic", "ll_two.*.two")
	rabbitmq.ConsumeTopic()
}
