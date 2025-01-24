package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hann10602/go-grpc/config"
)

func main() {
	ctx := context.Background()
	redisClient := config.NewRedisClient(ctx)
	channelName := fmt.Sprintf("notification: %s", "10602")

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down")
		case t := <-ticker.C:
			if cmd := redisClient.Publish(ctx, channelName, fmt.Sprintf("Take this: %s", t.String())); cmd.Err() != nil {
				panic(cmd.Err())
			}
		}
	}
}
