package main

import (
	"fmt"
	"rabbitMQ/pkg/rabbitMQ"
	"time"
)

func main() {
	rabbitmq1 := rabbitMQ.NewRabbitMQRouting("exnewPub", "ll_one")
	rabbitmq2 := rabbitMQ.NewRabbitMQRouting("exnewPub", "ll_two")
	for i := 0; i < 100; i++ {
		msgone := fmt.Sprintf("路由模式的消息 one Hello ll:%d", i)
		msgtwo := fmt.Sprintf("路由模式的消息 two Hello ll:%d", i)
		rabbitmq1.PublishRouting(msgone)
		rabbitmq2.PublishRouting(msgtwo)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("消息发送成功")

}
