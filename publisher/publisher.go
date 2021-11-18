package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
)

func Publish(ctx context.Context, topic *pubsub.Topic, content interface{}) {
	data, _ := json.Marshal(content)

	res := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, _ = res.Get(ctx)
}
