package usecase

import (
	"context"
	"time"

	"analytics-service/internal/domain/entity"
	"analytics-service/internal/domain/repository"
)

type AnalyticsUsecase struct {
	repo repository.EventRepo
}

func NewAnalyticsUsecase(repo repository.EventRepo) *AnalyticsUsecase {
	return &AnalyticsUsecase{repo: repo}
}

func (uc *AnalyticsUsecase) RecordEvent(ctx context.Context, e entity.Event) error {
	return uc.repo.Save(ctx, e)
}

func (uc *AnalyticsUsecase) GetUserDayEvents(ctx context.Context, userID string, day time.Time) ([]entity.Event, error) {
	return uc.repo.GetByUserAndDay(ctx, userID, day)
}
