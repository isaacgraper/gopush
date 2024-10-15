package main

import (
	"fmt"
	"time"

	"github.com/isaacgraper/gopush/event"
	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/sender"
	"github.com/isaacgraper/gopush/trigger"
)

func main() {
	e := &event.Event{
		Name:  "new Event",
		Usage: "test",
		// Modify the struct values by passing new values for it
		Notification: notification.Notification{
			Message: "Something",
			SendAt:  time.Now(),
		},
		// Trigger with the given context need to have a condition for triggering this event
		Trigger: trigger.Trigger{
			Ctx: func(sender *sender.Sender, notification *notification.Notification) (bool, error) {
				// Set your endpoint to send a websocket to this address
				// Can set a threshold for the socket
				sender.Set("localhost://8080", 0)
				// Pass the notification values
				sender.Send(notification)
				return true, nil
			},
		},
	} // Logic for triggering the event can be set here or in the Ctx above
	fmt.Println(e)

}
