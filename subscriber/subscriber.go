package subscriber

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func InitOne(ctx context.Context, sub *pubsub.Subscription) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Println("[SUB-ONE] Got message: ", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("failed to pull messages: %v", err)
	}
}

func InitTwo(ctx context.Context, sub *pubsub.Subscription) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Println("[SUB-TWO] Got message: ", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("failed to pull messages: %v", err)
	}
}
