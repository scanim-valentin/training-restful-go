package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

// GetConversation Retrieve conversation between two users from the database and toggle online status for this users
func GetConversation(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var groupid database.GroupID
	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &groupid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan parameter: ", err)
		return
	}
	messages, err := database.GetMessages(groupid)
	// Call to database
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodGet, err))

	// Parsing result
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}
}
