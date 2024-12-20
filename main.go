package main

import (
	"github.com/isaacgraper/gopush/event"
	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/sender"
	"github.com/isaacgraper/gopush/trigger"
)

func main() {
	e := &event.Event{
		Name:  "new Event",
		Usage: "test",
		// Pass the endpoint for the notification been sent to
		Sender: sender.Sender{
			Endpoint: ":8080",
			// Can pass a threshold for multiple events handling
			Threshold: 2,
		},
		// Modify the struct values by passing new values for it
		Notification: notification.Notification{
			Message: "An event has been triggered and it send a message",
		},
		// Trigger with the given context need to have a condition for triggering this event
		Trigger: trigger.Trigger{
			Ctx: func() (bool, error) {
				return true, nil
			},
		},
	}
	// Logic for triggering the event can be set here or in the Ctx above
	event.Run(e)

}
