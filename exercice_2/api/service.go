package main

import (
	"fmt"
	"log"
	"net/http"

	_ "service/router"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var Port string = "3001"

var counter int = 0

func add(w http.ResponseWriter, r *http.Request) {

	counter++
	// TODO mutex to deal with concurrency
	fmt.Fprintf(w, "%v", counter)
}

// Toggle online status for user
func toggleOnlineStatus(w http.ResponseWriter, r *http.Request) {

}

// Registers a user with a username and a password
func registerNewUser(w http.ResponseWriter, r *http.Request) {

}

// Registers a user with a username and a password
func retreiveConversation(w http.ResponseWriter, r *http.Request) {

}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/add/", add)

	//CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Starting server
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)
	log.Fatal(http.ListenAndServe(":"+Port, handler))
}
