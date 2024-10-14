package model

import "time"

type Event struct {
	ID        string
	EventType string
	CreatedAt time.Time
}

type EventService interface {
	Event(id string) (*Event, error)
	Events() ([]*Event, error)
	CreateEvent(e *Event) error
	DeleteEvent(id string) error
}
