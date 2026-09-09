package postgres

import (
	"context"

	"analytics-service/internal/domain/entity"
	"analytics-service/internal/domain/repository"
	"analytics-service/internal/infrastructure/postgres/models"

	"gorm.io/gorm"
)

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) repository.EventRepo {
	return &EventRepo{db: db}
}

func (r *EventRepo) Save(ctx context.Context, ent entity.Event) error {
	m := toModel(ent)
	return r.db.WithContext(ctx).Create(&m).Error
}

func toModel(ent entity.Event) models.Event {
	return models.Event{
		EventType:  string(ent.EventType),
		Status:     string(ent.Status),
		ErrorCode:  ent.ErrorCode,
		UserID:     ent.UserID,
		IP:         ent.IP,
		OccurredAt: ent.OccurredAt,
	}
}
