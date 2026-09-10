package usecase

import "analytics-service/internal/domain/entity"

type EventUsecase interface {
	Save()
	GetByUserId(id string) []entity.Event
}
