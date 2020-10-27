package mq

import (
	"carrotCloud/config"
	"log"

	"github.com/streadway/amqp"
)

// rabbitMQ连接对象
var conn *amqp.Connection

// channel 通过 connection 获得，主要通过channel进行消息的发布和接收
var channel *amqp.Channel

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

func init() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	if !config.AsyncTransferEnable {
		return
	}
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}

// initChannel ：初始化一个channel
func initChannel() bool {
	// 1.判断 channel 是否已经创建了，是则直接返回True
	if channel != nil {
		return true
	}
	// 2. 获得 rabbitMQ的一个连接
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// 3. 打开一个 channel， 用于消息的发布与接收等
	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// Publish : 发布消息
// 参数如下：
// 	exchange ：交换机，生产者发布消息时需要将信息投递到交换机上，需要进行指定交换机的名字
// 	routingKey ：需要指定 routingKey
// 	msg ：具体消息内容
// 返回如下：
// 	True，表示成功，false表示失败
func Publish(exchange, routingKey string, msg []byte) bool {
	// 打印信息
	log.Println(string(msg))

	// 1.判断 Channel 是否正常
	if !initChannel() {
		return false
	}
	// 2. 执行消息的发布操作
	if nil == channel.Publish(
		exchange, // 指定交换机，由参数传入
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, // 最新版rabbitMQ该参数不起作用
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		// 如果 channel.Publish 返回的 error 为nil则成功发送，否则返回 false 失败
		return true
	}
	return false
}
