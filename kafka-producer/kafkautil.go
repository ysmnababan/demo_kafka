package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewProducer() *kafkaProducer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // ✅ correct
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()
	return &kafkaProducer{producer: producer}
}

func (p *kafkaProducer) Terminate() {
	p.producer.Close()
}

func (p *kafkaProducer) SendMessage(topic, key, value string) error {
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,

			// remember to add the partition manually
			// use 'PartitionAny' for routing the key automatically
			// without it, partitions will be 0
			Partition: kafka.PartitionAny,
		},
		Value: []byte(value),
		Key:   []byte(key),
	}
	err := p.producer.Produce(&msg, nil)
	if err != nil {
		return err
	}
	p.producer.Flush(1 * 1000)
	return nil
}
