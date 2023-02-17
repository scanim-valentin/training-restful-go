// Package 'methods' regroups the different methods that either alter the chat history database based on http requests or internal calls from the API itself
package methods

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// DB management
var db_username string = "chatapi"
var db_password string = "chatapi"
var db_ip string = "localhost:5432"
var db_connect string = fmt.Sprintf("postgresql://%s:%s@%s/todos?sslmode=disable", db_username, db_password, db_ip)
var db *sql.DB = nil
var db_name string = "chatsystem"

/*
* Setup (connect to database)
 */
func Setup() {
	db, err := sql.Open("postgres", db_connect)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(fmt.Sprintf("ALTER %s INSER TABLE [IF NOT EXIST] \"user_status\" ( \"username\", \"status\" )", db_name))
	if err != nil {
		log.Fatal(err)
	}
}

// Toggle online status for user
func toggleOnlineStatus(username string) {

}

/*
* Routable methods
 */

// Registers a user with a username and a password
func Register(w http.ResponseWriter, r *http.Request) {

	// Unpacking request data
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data["form"])

	// TODO : check if user exists
	_, err = db.Exec(fmt.Sprintf("ALTER DATABASE %s INSER TABLE  %s_history", db_name))
	if err != nil {
		log.Fatal(err)
	}
}

// Login a user with a username and a password and toggle online status for this user
func Login(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Retrieve online user list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Retrieve conversation between two user from the database and toggle online status for this user
func GetConversation(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// Add a message to a conversation between two user from the database
func SendMessage(w http.ResponseWriter, r *http.Request) {
	//TODO
}
