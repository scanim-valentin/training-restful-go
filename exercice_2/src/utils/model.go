// Package 'model' regroups the different data structures used within the API
package utils

import (
	"net"
	"time"
)

// Unique user identifier
type UserID int

// Unique message identifier
type MessageID int

// This could change in the future
type MessageContent string

// A user has a unique UserID, a name (not unique), an IP address and a Port.
// IP and Port are also used to indicated the user's status (connected or not)
type User struct {
	ID   UserID
	Name string
	IP   net.IP
	Port int
}

// A message contains the source user, the destination user and a message body
type Message struct {
	ID                  MessageID
	Source, Destination UserID
	Body                MessageBody
}

// A message body contains the message content and the time at which it was sent
type MessageBody struct {
	Content MessageContent
	Time    time.Time
}

/*JSON-aimed structs*/

// Login response is returned to login and register HTTP requests
type LoginResponse struct {
	ID       UserID
	UserList []User
}

// Conversation response is returned to select HTTP requests
type Conversation struct {
	User, Other   UserID
	MessageBodies []MessageBody
}
