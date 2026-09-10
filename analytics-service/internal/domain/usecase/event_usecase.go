package usecase

import (
	"context"
	"time"

	"analytics-service/internal/domain/entity"
)

type EventUsecase interface {
	RecordEvent(ctx context.Context, e entity.Event) error
	GetUserDayEvents(ctx context.Context, userID string, day time.Time) ([]entity.Event, error)
}
