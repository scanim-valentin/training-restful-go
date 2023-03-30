package messages

import (
	"encoding/json"
	"log"
	"net/http"
	"service/database"
	"service/utils"
	"time"
)

// SendMessage Add a message to a conversation between two users from the database
func SendMessage(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var message database.Message
	err := decoder.Decode(&message)
	message.Time = message.Time.Truncate(time.Millisecond)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error when decoding message: ", err)
		return
	}
	// SQL Queries
	_, err = database.NewMessage(message)
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPost, err))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
	}

}
