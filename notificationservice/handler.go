package notificationservice

import (
	"fmt"
	"time"

	"github.com/hann10602/go-grpc/notificationservice/notificationproto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var _notificationService notificationproto.NotificationServiceServer = &Handler{}

type Handler struct {
	notificationproto.UnimplementedNotificationServiceServer
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{redisClient: redisClient}
}

func (h *Handler) GetNotifications(req *notificationproto.NotificationRequest, stream grpc.ServerStreamingServer[notificationproto.Notification]) error {
	pubsub := h.redisClient.Subscribe(stream.Context(), fmt.Sprintf("notification: %s", req.GetUserId()))

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-pubsub.Channel():
			if err := stream.Send(&notificationproto.Notification{
				UserId:    req.GetUserId(),
				Content:   fmt.Sprintf("New notification at %s: %s", time.Now().String(), msg.Payload),
				CreatedAt: time.Now().UnixMilli(),
			}); err != nil {
				return fmt.Errorf("failed to send notification: %w", err)
			}
		}
	}
}
