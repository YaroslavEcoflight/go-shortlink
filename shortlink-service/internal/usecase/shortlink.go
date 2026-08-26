package usecase

import (
	"shortlink-service/internal/domain/repository"
)

type ShortLinkService struct {
	repo repository.LinkRepo
}

func NewShortLinkService(repo repository.LinkRepo) ShortLinkService {
	return ShortLinkService{repo: repo}
}
