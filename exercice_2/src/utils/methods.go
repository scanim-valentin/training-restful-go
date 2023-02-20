// Package 'methods' regroups the different methods that either alter the chat history database based on http requests or internal calls from the API itself
package utils

import (
	"database/sql"
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

// Net complement
var ip_unspecified net.IP = net.IPv4(0, 0, 0, 0)

/*
* Setup (connect to database)
 */
func Setup() {
	var err error
	db, err = sql.Open("postgres", db_connect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to ", db_connect)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, username text, ip inet, port int )")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created or detected table 'user'")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, source integer, destination integer, content text )")
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

/*
* Routable methods
 */

// Registers a user with a username and save ip and port
func Register(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Registering new user with name ", values["name"][0])
	ip, port := getIP(r)
	fmt.Println("Extracted IP from request ", ip, port)
	// SQL Queries
	var id int
	err := db.QueryRow(fmt.Sprintf("INSERT INTO users (username, ip, port) VALUES ('%s', '%s', '%s') RETURNING id", values["name"][0], ip, port)).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Registered new user with name ", values["name"][0], "and ID ", id)
	fmt.Fprintf(w, "%v", id)
}

// Login a user with a username and save ip and port
func Login(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login in user with ID ", values["id"][0])
	ip, port := getIP(r)
	fmt.Println("Extracted IP from request ", ip, port)
	// SQL Queries
	_, err := db.Exec(fmt.Sprintf("UPDATE users SET ip = '%s', port = %s WHERE id = %s", ip, port, values["id"][0]))
	if err != nil {
		log.Fatal(err)
	}
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
func getIP(req *http.Request) (net.IP, string) {
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Panic("Error getIP: ", err)
	}
	return net.ParseIP(ip), port
}

// Logout a user: replaces ip and port by unspecified and 0
func Logout(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login out user with ID ", values["id"][0])
	// SQL Queries
	_, err := db.Exec(fmt.Sprintf("UPDATE users SET ip = '%s' WHERE id = %s", ip_unspecified, values["id"][0]))
	if err != nil {
		log.Fatal(err)
	}

}
