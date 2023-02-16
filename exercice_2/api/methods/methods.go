package methods

import (
	"fmt"
	"net/http"
)

// DB management
var DB_username string = ""
var DB_password string = ""
var DB_ip string = ""
var DB string = fmt.Sprintf("postgresql://%s:%s@%s/todos?sslmode=disable", DB_username, DB_password, DB_ip)

// Toggle online status for user
func toggleOnlineStatus(username string) {

}

/*
* Routable methods
 */

// Registers a user with a username and a password
func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Login a user with a username and a password and toggle online status for this user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Retrieve online user list
func RetrieveOnlineUserList(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Retrieve conversation between two user from the database and toggle online status for this user
func RetrieveConversation(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Add a message to a conversation between two user from the database
func AddMessageToConversation(w http.ResponseWriter, r *http.Request) {
	//TODO
}
