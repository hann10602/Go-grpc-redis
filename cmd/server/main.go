package main

import (
	"context"
	"log"
	"log/slog"
	"net"

	"github.com/hann10602/go-grpc/config"
	"github.com/hann10602/go-grpc/notificationservice"
	"github.com/hann10602/go-grpc/notificationservice/notificationproto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", notificationservice.Address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	ctx := context.Background()

	redisClient := config.NewRedisClient(ctx)
	handler := notificationservice.NewHandler(redisClient)
	notificationproto.RegisterNotificationServiceServer(grpcServer, handler)
	slog.Info("Listening on: " + notificationservice.Address)
	grpcServer.Serve(lis)
}
