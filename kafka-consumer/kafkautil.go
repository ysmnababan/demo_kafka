package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type messageBroker struct {
	consumer *kafka.Consumer
}

func NewKafkaUtil(consumerGroup string) *messageBroker {
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
	return &messageBroker{consumer: consumer}
}

func (m *messageBroker) Subscribe(topic string) {
	err := m.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
}

func (m *messageBroker) ReadMessage(timeout time.Duration) (message *kafka.Message, err error) {
	message, err = m.consumer.ReadMessage(timeout)
	return
}

func (m *messageBroker) Terminate() error {
	return m.consumer.Close()
}
