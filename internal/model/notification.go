package model

type Notification struct {
	ID        string
	UserID    string
	Message   string
	EventType string
	CreatedAt string
}

type NotificationService interface {
	Notification(id string) (*Notification, error)
	Notifications() ([]*Notification, error)
	CreateNotification(e *Notification) error
	DeleteEvent(id string) error
}
