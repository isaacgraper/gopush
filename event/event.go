package event

import "github.com/isaacgraper/gopush/trigger"

type Event struct {
	Name    string
	Usage   string
	Trigger trigger.Trigger
}

type iEvent interface {
}
