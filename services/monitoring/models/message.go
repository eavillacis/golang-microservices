package models

// DeadLetterMessage ...
type DeadLetterMessage struct {
	Data         string
	Subscription string
	Message      string
}
