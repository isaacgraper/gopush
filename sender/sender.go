package sender

import (
	"fmt"

	"github.com/isaacgraper/gopush/notification"
)



// endpoint for sending a socket trought
// threshold for socket waiting
type Sender struct {
	Endpoint  string
	Threshold float64
}

// Modify the valeus for posterior usage
func (s *Sender) Set(endpoint string, threshold float64) {
	s.Endpoint = endpoint
	s.Threshold = threshold
}

// Will be implemented the logic for sending the socket here and the given message
// Can be implemented more methods
func (s *Sender) Send(message *notification.Notification) bool {
	if s.Endpoint == "" {
		fmt.Println("Set not configured with an endpoint")
		return false
	}
	fmt.Printf("Sending \"%s\" to %s\n", message, s.Endpoint)
	return true
}
