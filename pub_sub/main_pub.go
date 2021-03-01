package main

import (
	"fmt"
	"rabbitMQ/pkg/rabbitMQ"
	"time"
)

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQPubSub("exNewPub")
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("路由模式的消息 Hello ll:%d", i)
		fmt.Println(msg)
		rabbitmq.PublishPub(msg)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("消息发送成功")

}
