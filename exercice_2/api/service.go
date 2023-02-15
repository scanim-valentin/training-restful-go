package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Port string = "3001"
var Dir http.Dir = http.Dir("./static")
var Prefix string = "/public/"

var counter int = 0

func add(w http.ResponseWriter, r *http.Request) {

	counter++
	// TODO mutex to deal with concurrency
	fmt.Fprintf(w, "%v", counter)
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
