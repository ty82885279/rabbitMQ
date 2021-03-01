package main

import "rabbitMQ/pkg/rabbitMQ"

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQRouting("exnewPub", "ll_one")
	rabbitmq.ConsumeRouting()
}
