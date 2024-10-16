package trigger

import (
	"errors"
	"time"

	"github.com/isaacgraper/gopush/sender"
)

type Cases interface {
	// Handle failture cases
	OnFailure(err error) error
	// Execute the trigger
	Execute() error
	// Validade input data
	Validate() bool
	// Set a timeout for the execution
	Timeout(timeout *sender.Sender) time.Duration
}

type Trigger struct {
	Ctx func() (bool, error)
}

func (t *Trigger) Execute() error {
	success, err := t.Ctx()
	if err != nil {
		return t.OnFailure(err)
	}
	if !success {
		return errors.New("trigger did not succeed")
	}
	return nil
}

func (t *Trigger) Validate(data any) bool {
	return data != nil
}

func (t *Trigger) OnFailure(err error) error {
	return errors.New("trigger failed: " + err.Error())
}
