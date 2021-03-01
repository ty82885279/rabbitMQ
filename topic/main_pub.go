package main

import (
	"fmt"
	"rabbitMQ/pkg/rabbitMQ"
	"time"
)

func main() {
	rabbitmq1 := rabbitMQ.NewRabbitMQTopic("exNewTopic", "ll_one.top.one")
	rabbitmq2 := rabbitMQ.NewRabbitMQTopic("exNewTopic", "ll_two.top.two")
	for i := 0; i < 100; i++ {
		msgone := fmt.Sprintf("Topic模式的消息 one Hello ll:%d", i)
		msgtwo := fmt.Sprintf("Topic模式的消息 two Hello ll:%d", i)
		rabbitmq1.PublishTopic(msgone)
		rabbitmq2.PublishTopic(msgtwo)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("消息发送成功")

}
