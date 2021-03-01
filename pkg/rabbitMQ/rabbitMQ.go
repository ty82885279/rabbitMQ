package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	con       *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string //交换机
	Key       string //Key
	MqURL     string //连接信息
}

// 创建rabbitMQ结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {

	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqURL:     MQURL,
	}
	var err error
	rabbitmq.con, err = amqp.Dial(rabbitmq.MqURL)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.con.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.con.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

//  简单模式Step:1.创建简单模式下的rabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

//  简单模式Step:2.简单模式下的生产消息
func (r *RabbitMQ) PublishSimple(msg string) {
	// 1.尝试申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true,找不到队列，会把消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列中，队列没有绑定消费者，会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

//  简单模式Step:3.简单模式下的消费消息
func (r *RabbitMQ) ConsumeSimple() {
	// 1.尝试申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    //用来区分多个消费者
		true,  //消费被消费后， 是否自动应答
		false, //是否具有排他性
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a msg:%s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for msg，To exit press Ctrl+C")
	select {}
}

// 订阅模式Step:1.创建rabbitMQ实例
func NewRabbitMQPubSub(exchange string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, "")
	return rabbitmq
}

// 订阅模式Step:2.生产消息
func (r *RabbitMQ) PublishPub(msg string) {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		"",
		// 如果为true,找不到队列，会把消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列中，队列没有绑定消费者，会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

// 订阅模式Step:3.消费消息
func (r *RabbitMQ) ConsumePub() {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.尝试创建队列，队列名称不要写，随机队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "创建队列失败")

	// 3.绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"", //pub/sub模式下,key为空
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(
		q.Name,
		"",    //用来区分多个消费者
		true,  //消费被消费后， 是否自动应答
		false, //是否具有排他性
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a msg:%s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for msg，To exit press Ctrl+C")
	select {}
}

// 路由模式Step:1.创建rabbitMQ实例
func NewRabbitMQRouting(exchange, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, routingKey)
	return rabbitmq
}

// 路由模式Step:2.生产消息
func (r *RabbitMQ) PublishRouting(msg string) {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		// 如果为true,找不到队列，会把消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列中，队列没有绑定消费者，会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

// 路由模式Step:3.消费消息
func (r *RabbitMQ) ConsumeRouting() {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.尝试创建队列，队列名称不要写，随机队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "创建队列失败")

	// 3.绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(
		q.Name,
		"",    //用来区分多个消费者
		true,  //消费被消费后， 是否自动应答
		false, //是否具有排他性
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a msg:%s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for msg，To exit press Ctrl+C")
	select {}
}

// Topic模式Step:1.创建rabbitMQ实例
func NewRabbitMQTopic(exchange, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, routingKey)
	return rabbitmq
}

// Topic模式Step:2.生产消息
func (r *RabbitMQ) PublishTopic(msg string) {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		// 如果为true,找不到队列，会把消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列中，队列没有绑定消费者，会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

// Topic模式Step:3.消费消息
func (r *RabbitMQ) ConsumeTopic() {
	// 1.尝试申请交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.尝试创建队列，队列名称不要写，随机队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "创建队列失败")

	// 3.绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(
		q.Name,
		"",    //用来区分多个消费者
		true,  //消费被消费后， 是否自动应答
		false, //是否具有排他性
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a msg:%s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for msg，To exit press Ctrl+C")
	select {}
}
