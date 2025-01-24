package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}

	return client
}
