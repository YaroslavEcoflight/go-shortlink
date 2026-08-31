package postgres

import (
	"shortlink-service/internal/domain/entity"
	"shortlink-service/internal/infrastructure/postgres/models"
	"shortlink-service/internal/pkg/base62"

	"gorm.io/gorm"
)

type LinkRepo struct {
	db *gorm.DB
}

func NewLinkRepo(db *gorm.DB) LinkRepo {
	return LinkRepo{db: db}
}

func (r *LinkRepo) Create(e entity.Link) (entity.Link, error) {
	m := toModel(e)
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&m).Error; err != nil {
			return err
		}
		m.Code = base62.Encode(m.ID)
		return tx.Save(&m).Error
	})
	if err != nil {
		return entity.Link{}, err
	}
	return toEntity(m), nil
}

func (r *LinkRepo) Delete(code string) error {
	return r.db.Where("code = ?", code).Delete(&models.Link{}).Error
}

func (r *LinkRepo) GetByCode(code string) (entity.Link, error) {
	var m models.Link
	if err := r.db.Where("code = ?", code).First(&m).Error; err != nil {
		return entity.Link{}, err
	}
	return toEntity(m), nil
}

func toModel(e entity.Link) models.Link {
	return models.Link{
		ID:   e.Id,
		Url:  e.Url,
		Code: e.Code,
	}
}

func toEntity(link models.Link) entity.Link {
	return entity.Link{
		Id:   link.ID,
		Code: link.Code,
		Url:  link.Url,
	}
}
