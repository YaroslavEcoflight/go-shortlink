package repository

import "shortlink-service/internal/domain/entity"

type LinkRepo interface {
	GetByCode(code string) (entity.Link, error)
	Create(entity.Link) (entity.Link, error)
	Delete(code string) error
}
