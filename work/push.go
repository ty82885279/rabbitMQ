package main

import (
	"fmt"
	"rabbitMQ/pkg/rabbitMQ"
	"time"
)

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQSimple("LLSimple")
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Hello ll:%d", i)
		fmt.Println(msg)
		rabbitmq.PublishSimple(msg)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("消息发送成功")

}
