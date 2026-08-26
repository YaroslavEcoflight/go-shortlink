package postgres

import "gorm.io/gorm"

type LinkRepo struct {
	db gorm.DB
}
