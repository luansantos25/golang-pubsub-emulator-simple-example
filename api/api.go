package api

import (
	"cloud.google.com/go/pubsub"
	"context"
	"golang-pubsub/publisher"
	"log"
	"net/http"
)

func Init(ctx context.Context, topic *pubsub.Topic, port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		publisher.Publish(ctx, topic, request.URL.Query()["content"])
		_, _ = writer.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil), nil)
}
