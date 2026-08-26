package usecase

import (
	"shortlink-service/internal/domain/entity"
	"shortlink-service/internal/domain/repository"
	"shortlink-service/internal/domain/service"
)

type LinkService struct {
	repo repository.LinkRepo
}

func NewLinkService(repo repository.LinkRepo) service.LinkService {
	return &LinkService{repo: repo}
}

func (svc *LinkService) Create(ent entity.Link) (entity.Link, error) {
	res, err := svc.repo.Create(ent)
	if err != nil {
		return entity.Link{}, err
	}
	return res, nil
}

func (svc *LinkService) Delete(code string) error {
	return nil
}

func (svc *LinkService) GetByCode(code string) (entity.Link, error) {
	return entity.Link{}, nil
}
