package main

import (
	"fmt"
)

type Producer interface {
	Terminate()
	SendMessage(topic, key, value string) error
}

func main() {
	var p Producer = NewProducer()
	defer p.Terminate()
	topic := "helloworld"
	for i := range 10 {
		key := fmt.Sprintf("key-%d", i*21)
		value := fmt.Sprintf("Message %d", i)
		err := p.SendMessage(topic, key, value)
		if err != nil {
			fmt.Println("error sending message", err)
			continue
		}
	}
}
