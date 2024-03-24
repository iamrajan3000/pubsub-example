package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	projectID  = "your-project-id"
	topicID    = "example-topic"
	subscriber = "example-subscriber"
)

func publishMessage(client *pubsub.Client, topicID string, msg string) {
	ctx := context.Background()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	_, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Printf("Message published: %s\n", msg)
}

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	// Publish a message every second
	for {
		publishMessage(client, topicID, "Hello, Pub/Sub!")
		time.Sleep(time.Second)
	}
}
