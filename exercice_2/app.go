package main

import (
	"log"
	"net/http"
)

var Port string = "8085"
var Dir http.Dir = http.Dir("./static")
var Prefix string = "/public/"

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:"+Port+Prefix+"index.html", http.StatusMovedPermanently)
}

func main() {
	fileServer := http.StripPrefix(Prefix, http.FileServer(Dir))
	http.Handle(Prefix, fileServer)
	// Default handler
	http.HandleFunc("/", index)

	// Starting the service
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
