package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaUtil(consumerGroup string) *kafkaConsumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          consumerGroup,
		"auto.offset.reset": "earliest",
		// "debug":             "consumer,cgrp,topic,broker",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("create new kafka")
	return &kafkaConsumer{consumer: consumer}
}

func (m *kafkaConsumer) Subscribe(topic string) {
	err := m.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
}

func (m *kafkaConsumer) ReadMessage(timeout time.Duration) (message *kafka.Message, err error) {
	message, err = m.consumer.ReadMessage(timeout)
	return
}

func (m *kafkaConsumer) Terminate() error {
	return m.consumer.Close()
}
