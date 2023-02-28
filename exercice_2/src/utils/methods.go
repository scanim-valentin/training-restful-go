// Package 'methods' regroups the different methods that either alter the chat history database based on http requests or internal calls from the API itself
package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/lib/pq"
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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, source integer, destination integer, content text, time timestamp )")
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

// Retrieve online user list
func getUserList() []User {
	// SQL Queries
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	// Reading rows
	for rows.Next() {
		var user User
		var ip_aux string
		if err := rows.Scan(&user.ID, &user.Name, &ip_aux, &user.Port); err != nil {
			return users
		}
		user.IP = net.ParseIP(ip_aux)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Panic(err)
	}
	return users
}

// https://blog.golang.org/context/userip/userip.go
func getIP(req *http.Request) (net.IP, string) {
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Panic("Error getIP: ", err)
	}
	return net.ParseIP(ip), port
}

/*
* Routable methods
 */

// Registers a user with a username and save ip and port
func Register(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Registering new user with name", values["name"][0])
	ip, port := getIP(r)
	fmt.Println("Extracted IP from request ", ip, port)

	// SQL Queries
	var id int
	err := db.QueryRow("INSERT INTO users (username, ip, port) VALUES ($1, $2, $3) RETURNING id", values["name"][0], ip.String(), fmt.Sprint(port)).Scan(&id)
	if err != nil {
		log.Panic(err)
	}

	// Parsing result
	fmt.Println("Registered new user with name ", values["name"][0], "and ID ", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{UserID(id), getUserList()})
}

// Login a user with a username and save ip and port
func Login(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login in user with ID ", values["id"][0])
	ip, port := getIP(r)
	fmt.Println("Extracted IP from request ", ip, port)
	// SQL Queries
	_, err := db.Exec("UPDATE users SET ip = $1, port = $2 WHERE id = $3", ip.String(), fmt.Sprint(port), values["id"][0])
	if err != nil {
		log.Fatal(err)
	}
	// Retrieving user list
	users := getUserList()
	for _, user := range users {
		fmt.Print(user)
	}
	// Parsing result
	w.Header().Set("Content-Type", "application/json")
	var id UserID
	if _, err := fmt.Sscanf(values["id"][0], "%d", &id); err != nil {
		log.Panic(err)
	}
	json.NewEncoder(w).Encode(LoginResponse{id, getUserList()})
}

// Retrieve conversation between two user from the database and toggle online status for this user
func GetConversation(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	userstr, otherstr := values["user"][0], values["other"][0]
	fmt.Print(userstr, otherstr)
	// SQL Queries
	rows, err := db.Query("SELECT * FROM messages WHERE source = $1 AND destination = $2 OR source = $2 AND destination = $1 ORDER BY time ASC", userstr, otherstr)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var messages []Message

	// Reading rows
	for rows.Next() {
		var message Message
		fmt.Print(rows)
		if err := rows.Scan(&message.ID, &message.Source, &message.Destination, &message.Content, &message.Time); err != nil {
			break
		}
		messages = append(messages, message)
	}
	fmt.Print(messages)
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	// Parsing result
	if err = json.NewEncoder(w).Encode(Conversation{messages}); err != nil {
		log.Panic(err)
	}
	fmt.Print(Conversation{messages})
}

// Add a message to a conversation between two user from the database
func SendMessage(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	decoder := json.NewDecoder(r.Body)

	var message Message
	err := decoder.Decode(&message)

	if err != nil {
		log.Panic(err)
	}
	fmt.Print(message)
	// SQL Queries
	var id MessageID
	err = db.QueryRow("INSERT INTO messages (source, destination, content, time) VALUES ($1, $2, $3, $4) RETURNING id",
		fmt.Sprint(message.Source), fmt.Sprint(message.Destination), message.Content, string(pq.FormatTimestamp(message.Time))).Scan(&id)
	if err != nil {
		// Message was not created
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		log.Print(err)
	} else {
		// Success in creating new message
		w.WriteHeader(http.StatusCreated)
		// Parsing result
		fmt.Fprintf(w, "%v", id)
	}
}

// Logout a user: replaces ip and port by unspecified and 0
func Logout(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login out user with ID ", values["id"][0])
	// SQL Queries
	_, err := db.Exec("UPDATE users SET ip = $1 WHERE id = $2", ip_unspecified.String(), values["id"][0])
	if err != nil {
		log.Fatal(err)
	}

}
