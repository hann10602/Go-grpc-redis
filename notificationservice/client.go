package notificationservice

import (
	"fmt"

	"github.com/hann10602/go-grpc/notificationservice/notificationproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (notificationproto.NotificationServiceClient, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(Address, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to construct client: %w", err)
	}

	return notificationproto.NewNotificationServiceClient(conn), nil
}
