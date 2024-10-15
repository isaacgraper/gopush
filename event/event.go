package event

import (
	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/trigger"
)

// The given struct event will be placed in a CLI tool for better debugging
type Event struct {
	Name         string
	Usage        string
	Notification notification.Notification
	Trigger      trigger.Trigger
}
