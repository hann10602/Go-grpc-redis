package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/hann10602/go-grpc/notificationservice"
	"github.com/hann10602/go-grpc/notificationservice/notificationproto"
)

func main() {
	client, err := notificationservice.NewClient()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	stream, err := client.GetNotifications(ctx, &notificationproto.NotificationRequest{UserId: "10602"})

	if err != nil {
		panic(err)
	}

	for {
		notification, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to read notification: %w", err)
		}

		b, err := json.MarshalIndent(notification, "", "\t")

		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(b))
	}
}
