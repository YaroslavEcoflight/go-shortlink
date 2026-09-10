package analytics

import (
	"context"
	"log"
	pb "shortlink-service/proto/analytics"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	pb pb.AnalyticsServiceClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{pb: pb.NewAnalyticsServiceClient(conn)}, nil
}

func (c *Client) RecordEvent(eventType string, status pb.EventStatus, userID, ip, errorCode string) {
	go func() {
		_, err := c.pb.RecordEvent(context.Background(), &pb.RecordEventRequest{
			EventType:  eventType,
			Status:     status,
			ErrorCode:  errorCode,
			UserId:     userID,
			Ip:         ip,
			OccurredAt: timestamppb.Now(),
		})
		if err != nil {
			log.Printf("analytics: failed to record event: %v", err)
		}
	}()
}
