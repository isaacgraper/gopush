package event

import (
	"fmt"

	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/sender"
	"github.com/isaacgraper/gopush/trigger"
)

// The given struct Event, the values will be placed in a CLI tool
type Event struct {
	Name         string
	Usage        string
	Sender       sender.Sender
	Notification notification.Notification
	Trigger      trigger.Trigger
}

func Run(e *Event) {
	fmt.Printf("Running event: %s\n", e.Name)
	fmt.Printf("Notification message: %s\n", e.Notification.Message)
	fmt.Printf("Hold for %d\n", e.Sender.Threshold)
	fmt.Printf("Sending to: %s\n", e.Sender.Endpoint)

	if e.Trigger.Ctx != nil {
		if _, err := e.Trigger.Ctx(); err != nil {
			fmt.Println("Error executing trigger", err)
		}

		e.Sender.Send(&e.Notification)
	}
}
