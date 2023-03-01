package main

import (
	"fmt"
	"log"
	"net/http"
	"service/database"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	router "service/server"
)

var Port string = "3001"

func main() {
	// Setting up routes
	router.Setup()

	//Setting up database
	database.Setup()
	defer database.Close()

	//CORS
	headersOk := handlers.AllowedHeaders([]string{"Content-Types"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Starting server
	fmt.Println("Starting server on port ", Port)
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(router.APIRouter)
	log.Fatal(http.ListenAndServe(":"+Port, handler))
}
