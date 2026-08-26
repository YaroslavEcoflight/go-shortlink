package service

import "shortlink-service/internal/domain/entity"

type LinkService interface {
	Create(entity entity.Link) (entity.Link, error)
	Delete(code string) error
	GetByCode(code string) (entity.Link, error)
}
