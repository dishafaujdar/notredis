package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"context"
)

func producer() {
	writer := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		Topic: "write_blog",
		Balancer: &kafka.Hash{},
		RequiredAcks: kafka.RequireAll,
		Async: false,
	}
	defer writer.Close()

	message := []kafka.Message{
		{
			Key:   []byte("user-123"),
			Value: []byte("My first blog"),
		},
		{
			Key:   []byte("user-124"),
			Value: []byte("My second blog"),
		},
		{
			Key:   []byte("user-125"),
			Value: []byte("My third blog"),
		},
	}

	err := writer.WriteMessages(context.Background(), message...)
	if err != nil {
		fmt.Println("Failed to write message: ", err)
		return
	}
	fmt.Println("Message written successfully !!")
}

func consumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "write_blog",
		GroupID: "blog_consumer",
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Failed to read message: ", err)
			continue
		}
		
		blog := string(message.Value) + "from kafka"

		saveBlog := func() {
			fmt.Println("Saving blog: ", blog)
		}
		go saveBlog()

		updateBlogCount :=func() {
				fmt.Println("Updating blog count: ", message.Key)
			}
		go updateBlogCount()
		
		fmt.Println("Blog saved and count updated successfully !!")

	}

}

func main() {
	go producer()
	consumer()
}