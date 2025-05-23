package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("here")
	k := NewKafkaUtil("golang")
	k.Subscribe("helloworld")
	defer func() {
		_ = k.Terminate()
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("here")
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Println("exit program:", sig)
			run = false
		default:
			message, err := k.ReadMessage(time.Second)
			if err != nil {
				continue
			}
			fmt.Printf("Consume event from topic: %s: partition: %v key = %s value = %s",
				*message.TopicPartition.Topic,
				message.TopicPartition.Partition,
				message.Key,
				message.Value)
		}
	}
}
