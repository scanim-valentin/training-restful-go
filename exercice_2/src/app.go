package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"service/server"
	// "fmt"
	// "log"
	// "net/http"
	"service/database"
	// "github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

var Port string = "3001"
var IP string = "0.0.0.0"

func main() {
	// Setting up routes
	server.Setup()

	//Setting up database
	database.Setup()
	defer database.Close()
	//CORS
	headersOk := handlers.AllowedHeaders([]string{"Content-Types"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Starting server
	fmt.Println("Starting server on port ", Port)
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(server.APIRouter)
	log.Fatal(http.ListenAndServe(IP+":"+Port, handler))
}
