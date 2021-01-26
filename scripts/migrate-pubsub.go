package main

import (	
	"context"
	"log"

	"github.com/eavillacis/velociraptor/pkg/subscriptions"
	"github.com/eavillacis/velociraptor/pkg/topics"
	"cloud.google.com/go/pubsub"
)

func createTopicIfNotExists(c *pubsub.Client, topic string) *pubsub.Topic {
	ctx := context.Background()

	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}

	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}

	return t
}

func createSubscriptionIfNotExists(c *pubsub.Client, topic, subscription string) *pubsub.Subscription {
	ctx := context.Background()

	s := c.Subscription(subscription)
	ok, err := s.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return s
	}

	t := c.Topic(topic)
	s, err = c.CreateSubscription(ctx, subscription, pubsub.SubscriptionConfig{Topic: t})
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}

	return s
}

func main() {
	ctx := context.Background()

	// For now the projectID is hard-coded
	pubsubClient, err := pubsub.NewClient(ctx, "digital-platforms-302300")
	if err != nil {
		log.Fatalf("Failed to generate a new pubsub client: %+v", err)
	}

	createTopicIfNotExists(pubsubClient, topics.DeadLetters)
	createTopicIfNotExists(pubsubClient, topics.Monitoring)

	createSubscriptionIfNotExists(pubsubClient, topics.DeadLetters, subscriptions.DeadLetterMonitoring)
	createSubscriptionIfNotExists(pubsubClient, topics.Monitoring, subscriptions.Monitoring)
}
