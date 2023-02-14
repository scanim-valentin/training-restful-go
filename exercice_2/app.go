package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	// Handles top-level page.
	content, err := ioutil.ReadFile("html/index.html")
	if err != nil {
		fmt.Fprintf(writer, "error: %s", err)
		log.Fatal(err)
	} else {
		fmt.Fprint(writer, string(content))
	}
}

func main() {
	// Default handler
	http.HandleFunc("/", index)

	// Starting the service
	log.Fatal(http.ListenAndServe(":8080", nil))
}
