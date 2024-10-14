package trigger

import "github.com/isaacgraper/gopush/notification"

type Trigger struct {
	Ctx          func() bool
	Threshold    float64
	Notification notification.Notification
}
