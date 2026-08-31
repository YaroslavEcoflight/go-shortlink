package models

import "time"

type Link struct {
	ID        int64  `gorm:"primaryKey"`
	Code      string `gorm:"unique;not null"`
	Url       string `gorm:"not null"`
	CreateAt  time.Time
	UpdatedAt time.Time
}

func (Link) TableName() string {
	return "links"
}
