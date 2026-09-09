package models

import "time"

type Event struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	EventType  string `gorm:"not null;index"`
	Status     string `gorm:"not null;index"`
	ErrorCode  string
	UserID     string    `gorm:"index"`
	IP         string    `gorm:"not null"`
	OccurredAt time.Time `gorm:"not null;index"`
}

func (Event) TableName() string {
	return "events"
}
