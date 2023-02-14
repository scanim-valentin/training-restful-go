package main

import (
	"log"
	"net/http"
)

var Port string = "8085"
var Dir http.Dir

// var err error

/*
func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/static/index.html", http.StatusMovedPermanently)
}
*/

func main() {
	Dir = http.Dir("./static")
	fileServer := http.FileServer(Dir)
	http.Handle("/test", fileServer)
	// Default handler
	// http.HandleFunc("/", index)

	// Starting the service
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
