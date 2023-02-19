// Package 'methods' regroups the different methods that either alter the chat history database based on http requests or internal calls from the API itself
package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"
)

// DB management
var db_username string = "postgres"
var db_password string = "PassWord"
var db_ip string = "localhost:5432"
var db_connect string = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", db_username, db_password, db_ip, db_name)
var db *sql.DB = nil
var db_name string = "postgres"

/*
* Setup (connect to database)
 */
func Setup() {
	db, err := sql.Open("postgres", db_connect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to ", db_connect)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users ( username text, ip cidr, port smallint )")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created or detected table 'user'")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages ( source integer, destination integer, content text )")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created or detected table 'messages'")
}

/*
* Should be called before closing service
 */
func Close() {
	log.Fatal(db.Close())
	fmt.Println("Successfully disconnected from ", db_connect)
}

// Toggle online status for user
func toggleOnlineStatus(username string) {

}

/*
* Routable methods
 */

// Registers a user with a username and save ip and port
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

// Login a user with a username and save ip and port
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

// https://blog.golang.org/context/userip/userip.go
func getIP(req *http.Request) (net.IP, int) {
	ip, port_str, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Printf("Error getIP: ", err)
	}
	var port int
	_, err = fmt.Sscanf(port_str, "%d", &port)
	if err != nil {
		log.Fatal(err)
	}
	return net.ParseIP(ip), port
}

// Logout a user: replaces ip and port by unspecified and 0
func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO
}
