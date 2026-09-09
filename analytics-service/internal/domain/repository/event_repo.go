package repository

import (
	"analytics-service/internal/domain/entity"
	"context"
)

type EventRepo interface {
	Save(ctx context.Context, event entity.Event) error
}
