package notificationservice

import "github.com/redis/go-redis/v9"

type Handler struct {
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{redisClient: redisClient}
}