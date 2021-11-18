package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"golang-pubsub/api"
	"golang-pubsub/subscriber"
	"google.golang.org/api/option"
	"log"
	"os"
	"sync"
)

const (
	projectIDEnv     = "PROJECT_ID"
	topicIDEnv       = "TOPIC_ID"
	oneSubIDEnv      = "ONE_SUBSCRIPTION_ID"
	twoSubIDEnv      = "TWO_SUBSCRIPTION_ID"
	apiServerPortEnv = "API_SERVER_PORT"
)

func main() {
	var (
		projectID     = os.Getenv(projectIDEnv)
		topicID       = os.Getenv(topicIDEnv)
		oneSubID      = os.Getenv(oneSubIDEnv)
		twoSubID      = os.Getenv(twoSubIDEnv)
		apiServerPort = os.Getenv(apiServerPortEnv)
	)

	ctx := context.Background()

	client := newPubSubClient(ctx, projectID)
	defer client.Close()

	topic := getOrCreateTopic(ctx, client, topicID)

	cfg := &pubsub.SubscriptionConfig{
		Topic: topic,
	}

	subOne := getOrCreateSub(ctx, client, topicID, oneSubID, cfg)
	subTwo := getOrCreateSub(ctx, client, topicID, twoSubID, cfg)

	var wg sync.WaitGroup

	wg.Add(3)
	go api.Init(ctx, topic, apiServerPort)
	go subscriber.InitOne(ctx, subOne)
	go subscriber.InitTwo(ctx, subTwo)
	wg.Wait()
}

func newPubSubClient(ctx context.Context, projectID string) *pubsub.Client {
	opts := []option.ClientOption{option.WithoutAuthentication()}

	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		log.Fatal("error when trying to create pubsub client", err)
	}

	return client
}

func getOrCreateTopic(ctx context.Context, client *pubsub.Client, topicID string) *pubsub.Topic {
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal("error when trying to check topic")
	}
	if !exists {
		_, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return topic
}

func getOrCreateSub(ctx context.Context, client *pubsub.Client, topicID, subID string, cfg *pubsub.SubscriptionConfig) *pubsub.Subscription {
	sub := client.Subscription(subID)
	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatalf("failed to check if subscription exists: %v", err)
	}
	if !ok {
		sub, err = client.CreateSubscription(ctx, subID, *cfg)
		if err != nil {
			log.Fatalf("failed to create subscription (%q): %v", topicID, err)
		}
	}
	return sub
}
