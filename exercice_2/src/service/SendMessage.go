package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
)

// SendMessage Add a message to a conversation between two user from the database
func SendMessage(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	var message database.Message
	err := decoder.Decode(&message)

	if err != nil {
		log.Panic(err)
	}
	fmt.Print(message)
	// SQL Queries
	_, err = database.NewMessage(message.Source, message.Destination, message.Content, message.Time)
	if err != nil {
		// Message was not created
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			return
		}
		log.Print(err)
	} else {
		// Success in creating new message
		w.WriteHeader(http.StatusCreated)
	}
}
