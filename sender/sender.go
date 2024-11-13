package sender

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/isaacgraper/gopush/notification"
)

type Action interface {
	Set(*Sender) error
	Send(endpoint string, threshold int) error
}

type Sender struct {
	Endpoint  string // Endpoint for sending a socket trought
	Threshold int    // Threshold for socket waiting
	// Must be implemented a time to live field
	// Must be implemented a semaphore condition for requests concurrency
}

// Set the values for different endpoints and threshold requests
func (s *Sender) Set(endpoint string, threshold int) {
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
	fmt.Printf("Sending: \"%s\" to %s at %s\n", message.Message, s.Endpoint, time.Now().Format("2006/01/02 15:04:05"))

	http.HandleFunc("/", s.wsHandler)

	err := http.ListenAndServe(s.Endpoint, nil)
	if err != nil {
		log.Fatal("Failed initializing server")
	}

	return true
}

func (s *Sender) wsHandler(w http.ResponseWriter, r *http.Request) {
	con, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify:   true,
		CompressionThreshold: s.Threshold,
	})
	if err != nil {
		log.Fatal("Failed to accept websocket connection: ", err)
	}

	defer con.Close(websocket.StatusNormalClosure, "Conexão encerrada")

	_, data, err := con.Read(r.Context())
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
		log.Println("Client closed connection")
		return
	} else if err != nil {
		log.Println("Error while trying to read data from the client:", err)
		return
	}
	log.Println("Data received from websocket:", string(data))
}
