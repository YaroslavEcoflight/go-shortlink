package usecase

import (
	"shortlink-service/internal/domain/entity"
	"shortlink-service/internal/domain/repository"
	"shortlink-service/internal/domain/service"
)

type LinkService struct {
	repo  repository.LinkRepo
	cache repository.LinkCache
}

func NewLinkService(repo repository.LinkRepo, cache repository.LinkCache) service.LinkService {
	return &LinkService{repo: repo, cache: cache}
}

func (svc *LinkService) Create(link entity.Link) (entity.Link, error) {
	res, err := svc.repo.Create(link)
	if err != nil {
		return entity.Link{}, err
	}
	_ = svc.cache.Set(res.Code, res)
	return res, nil
}

func (svc *LinkService) Delete(code string) error {
	if err := svc.repo.Delete(code); err != nil {
		return err
	}
	_ = svc.cache.Delete(code)
	return nil
}

func (svc *LinkService) GetByCode(code string) (entity.Link, error) {
	if link, err := svc.cache.Get(code); err == nil {
		return link, nil
	}

	link, err := svc.repo.GetByCode(code)
	if err != nil {
		return entity.Link{}, err
	}

	_ = svc.cache.Set(code, link)
	return link, nil
}
