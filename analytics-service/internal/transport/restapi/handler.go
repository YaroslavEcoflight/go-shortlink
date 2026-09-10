package restapi

import "analytics-service/internal/domain/usecase"

type AnalyticsEventHandler struct {
	svc usecase.EventUsecase
}

func NewHandler()
