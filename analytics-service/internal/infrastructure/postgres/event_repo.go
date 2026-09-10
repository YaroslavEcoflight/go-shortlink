package postgres

import (
	"context"
	"time"

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

func (r *EventRepo) GetByUserAndDay(ctx context.Context, userID string, day time.Time) ([]entity.Event, error) {
	var ms []models.Event
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := start.Add(24 * time.Hour)

	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND occurred_at >= ? AND occurred_at < ?", userID, start, end).
		Order("occurred_at asc").
		Find(&ms).Error; err != nil {
		return nil, err
	}

	events := make([]entity.Event, len(ms))
	for i, m := range ms {
		events[i] = toEntity(m)
	}
	return events, nil
}

func toEntity(m models.Event) entity.Event {
	return entity.Event{
		ID:         m.ID,
		EventType:  entity.EventType(m.EventType),
		Status:     entity.EventStatus(m.Status),
		ErrorCode:  m.ErrorCode,
		UserID:     m.UserID,
		IP:         m.IP,
		OccurredAt: m.OccurredAt,
	}
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
