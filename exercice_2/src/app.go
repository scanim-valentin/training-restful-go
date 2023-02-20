package main

import (
	"fmt"
	"log"
	"net/http"

	router "service/server"
	"service/utils"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

var Port string = "3001"

var counter int = 0

func add(w http.ResponseWriter, r *http.Request) {

	counter++
	// TODO mutex to deal with concurrency
	fmt.Fprintf(w, "%v", counter)
}

func main() {
	// Setting up routes
	router.Setup()

	//Setting up database
	utils.Setup()
	defer utils.Close()

	//CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Starting server
	fmt.Println("Starting server on port ", Port)
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(router.APIRouter)
	log.Fatal(http.ListenAndServe(":"+Port, handler))
}
