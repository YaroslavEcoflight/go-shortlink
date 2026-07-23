package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
