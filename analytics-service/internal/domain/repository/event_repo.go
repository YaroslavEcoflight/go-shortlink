package repository

import (
	"context"
	"time"

	"analytics-service/internal/domain/entity"
)

type EventRepo interface {
	Save(ctx context.Context, event entity.Event) error
	GetByUserAndDay(ctx context.Context, userID string, day time.Time) ([]entity.Event, error)
}
