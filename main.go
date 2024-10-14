package main

import (
	"log"
	"time"

	"github.com/isaacgraper/gopush/event"
	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/trigger"
)

func main() {
	e := &event.Event{
		Name:  "new Event",
		Usage: "test",
		Trigger: trigger.Trigger{
			Ctx: func() bool {
				return true
			},
			Threshold: 0,
			Notification: notification.Notification{
				Message: "send message",
				SendAt:  time.Unix(),
			},
		},
	}

	log.Println(e)

}
