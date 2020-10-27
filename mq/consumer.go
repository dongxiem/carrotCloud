package mq

import "log"

var done chan bool

// StartConsume : 开始监听队列，接收消息
// 参数如下：
// 	qName ：指定的队列名
// 	cName ：消费者的名称
// 	callback ：外部调用着指定的处理消息的回调函数
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	// 1.通过 channel.consume 获得消息信道
	msgs, err := channel.Consume(
		qName,
		cName,
		true,  // 自动应答
		false, // 非唯一的消费者，会有好几个消费者同时监听一个队列，需要根据竞争机制来进行
		false, // rabbitMQ只能设置为false
		false, // noWait, false表示会阻塞直到有消息过来
		nil)
	// 错误处理
	if err != nil {
		log.Fatal(err)
		return
	}

	done = make(chan bool)
	// 2. 开启一个 goroutine 循环获取队列的消息
	go func() {
		// 循环读取channel的数据
		for d := range msgs {
			// 3. 调用 callback 方法来处理新的消息
			processErr := callback(d.Body)
			if processErr {
				// TODO: 将任务写入错误队列，待后续处理，用于异常情况的重试
			}
		}
	}()

	// 接收done的信号, 没有信息过来则会一直阻塞，避免该函数退出
	<-done

	// 关闭通道
	channel.Close()
}

// StopConsume : 停止监听队列
func StopConsume() {
	done <- true
}
