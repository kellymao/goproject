package main

import (

	"time"

	"github.com/Shopify/sarama"
)
import "github.com/astaxie/beego/logs"



func send_to_kafka(topic string , msgtxt string) {

	// 往kafka里面发数据

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分布式，写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {

		logs.Error("producer close, err:", err)
		return
	}

	defer client.Close()

		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder(msgtxt)

		pid, offset, err := client.SendMessage(msg) // pid 是分区号，offset是分区里面的偏移量，
		if err != nil {
			logs.Error("send message failed,", err)
			return
		}

		logs.Debug("pid:%v offset:%v\n", pid, offset)
		time.Sleep(2 * time.Second)

}
