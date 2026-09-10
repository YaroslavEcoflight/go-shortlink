package grpc

import (
	"context"

	"analytics-service/internal/domain"
	"analytics-service/internal/domain/entity"
	"analytics-service/internal/usecase"
	pb "analytics-service/proto/analytics"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedAnalyticsServiceServer
	uc  *usecase.AnalyticsUsecase
	log domain.Interface
}

func NewHandler(uc *usecase.AnalyticsUsecase, log domain.Interface) *Handler {
	return &Handler{uc: uc, log: log}
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
	if err := h.uc.RecordEvent(ctx, e); err != nil {
		h.log.Error("RecordEvent type=%s user=%s: %v", req.EventType, req.UserId, err)
		return &emptypb.Empty{}, err
	}
	h.log.Debug("RecordEvent type=%s user=%s status=%s", req.EventType, req.UserId, req.Status)
	return &emptypb.Empty{}, nil
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
