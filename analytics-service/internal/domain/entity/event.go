package entity

import "time"

type EventType string
type EventStatus string

const (
	EventRegister EventType = "register"
	EventLogin    EventType = "login"
	EventLogout   EventType = "logout"
	EventRefresh  EventType = "refresh"
	EventValidate EventType = "validate"
)

const (
	StatusSuccess EventStatus = "success"
	StatusFailure EventStatus = "failure"
)

type Event struct {
	ID         int64
	EventType  EventType
	Status     EventStatus
	ErrorCode  string
	UserID     string
	IP         string
	OccurredAt time.Time
}
