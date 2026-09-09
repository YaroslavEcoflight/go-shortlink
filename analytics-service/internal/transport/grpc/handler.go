package grpc

import (
	"context"

	"analytics-service/internal/domain/entity"
	"analytics-service/internal/usecase"
	pb "analytics-service/proto/analytics"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedAnalyticsServiceServer
	uc *usecase.AnalyticsUsecase
}

func NewHandler(uc *usecase.AnalyticsUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) RecordEvent(ctx context.Context, req *pb.RecordEventRequest) (*emptypb.Empty, error) {
	e := entity.Event{
		EventType:  entity.EventType(req.EventType),
		Status:     protoStatusToEntity(req.Status),
		ErrorCode:  req.ErrorCode,
		UserID:     req.UserId,
		IP:         req.Ip,
		OccurredAt: req.OccurredAt.AsTime(),
	}
	return &emptypb.Empty{}, h.uc.RecordEvent(ctx, e)
}

func protoStatusToEntity(s pb.EventStatus) entity.EventStatus {
	switch s {
	case pb.EventStatus_STATUS_SUCCESS:
		return entity.StatusSuccess
	case pb.EventStatus_STATUS_FAILURE:
		return entity.StatusFailure
	default:
		return entity.StatusSuccess
	}
}
