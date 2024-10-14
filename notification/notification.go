package notification

import "time"

type Notification struct {
	Message string
	SendAt  time.Time
}
