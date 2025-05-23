package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer interface {
	Subscribe(topic string)
	ReadMessage(timeout time.Duration) (message *kafka.Message, err error)
	Terminate() error
}

func main() {
	var c Consumer = NewKafkaUtil("golang")
	c.Subscribe("helloworld")
	defer func() {
		_ = c.Terminate()
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Println("exit program:", sig)
			run = false
		default:
			message, err := c.ReadMessage(time.Second)
			if err != nil {
				continue
			}
			fmt.Printf("Consume event from topic: %s: partition: %v key = %s value = %s\n",
				*message.TopicPartition.Topic,
				message.TopicPartition.Partition,
				message.Key,
				message.Value)
		}
	}
}
