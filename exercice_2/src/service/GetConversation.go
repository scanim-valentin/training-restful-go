package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
)

// GetConversation Retrieve conversation between two user from the database and toggle online status for this user
func GetConversation(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var user, other database.UserID
	if _, err := fmt.Sscanf(values["user"][0], "%d", &user); err != nil {
		log.Panic(err)
	}
	if _, err := fmt.Sscanf(values["other"][0], "%d", &other); err != nil {
		log.Panic(err)
	}
	messages := database.GetMessages(user, other)
	// fmt.Println(messages)

	// Parsing result
	if err := json.NewEncoder(w).Encode(database.Conversation{Messages: messages}); err != nil {
		log.Panic(err)
	}
}
