package main

import (
	"fmt"
	"rabbitMQ/pkg/rabbitMQ"
)

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQSimple("LLSimple")
	rabbitmq.PublishSimple("Hello ll~")
	fmt.Println("消息发送成功")

}
