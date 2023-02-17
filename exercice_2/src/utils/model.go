// Package 'model' regroups the different data structures used within the API
package main

import "fmt"

type UserID int64
type MessageBody string

// A user has a unique UserID, a name (not unique), and a status which indicates whether or not the user is connected
type User struct {
	id     UserID
	name   string
	status bool
}

// A message contains the source user, the destination user and a body
type Message struct {
	source      UserID
	destination UserID
	body        *MessageBody
}

/*
// Returns a pointer to an allocated and initialised User instance
func UserInit(id UserID, name string, status bool) *User {
	return &User{id, name, status}
}

// Returns a pointer to an allocated and initialised Message instance
func MessageInit(source UserID, destination UserID, content *MessageBody) *Message {
	return &Message{source, destination, content}
}
*/

func main() {
	fmt.Printf("%v", *(UserInit(5585, "Tricotin", true)))
	fmt.Printf("%v", *(MessageInit(5585, 5586, "Saperlotte")))
}
