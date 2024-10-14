package model

import "time"

type Event struct {
	ID        string
	EventType string
	CreatedAt time.Time
}
