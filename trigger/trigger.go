package trigger

import (
	"github.com/isaacgraper/gopush/notification"
	"github.com/isaacgraper/gopush/sender"
)

// Need to be a bool for trigger the event to send for any occasion and use case
type Trigger struct {
	Ctx func(sender *sender.Sender, notification *notification.Notification) (bool, error)
}
