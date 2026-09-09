package grpc

import (
	pb "analytics-service/proto/analytics"

	"google.golang.org/grpc"
)

func NewServer(h *Handler) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterAnalyticsServiceServer(s, h)
	return s
}
