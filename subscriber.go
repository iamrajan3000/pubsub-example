package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
)

func handleMessage(msg *pubsub.Message) {
	fmt.Printf("Received message: %s\n", string(msg.Data))
	msg.Ack()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subscriber)
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handleMessage(msg)
	})
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}

	// Handle termination signals to gracefully stop the subscriber
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Shutting down subscriber...")
	cancel()
	time.Sleep(time.Second) // Wait for the context to cancel
	fmt.Println("Subscriber stopped.")
}
