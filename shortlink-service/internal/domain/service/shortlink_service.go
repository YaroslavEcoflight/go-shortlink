package service

import "shortlink-service/internal/domain/entity"

type LinkService interface {
	CreateLink(entity.Link) error
	DeleteLink(code string) error
	GetLinkByCode(code string) entity.Link
}
