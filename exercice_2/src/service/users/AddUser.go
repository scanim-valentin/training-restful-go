package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

// AddUser Registers a users with a username and save ip and port
func AddUser(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var user database.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Error when decoding user: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SQL Queries
	id, err := database.NewUser(user)
	if err != nil {
		log.Println("Error when creating new user: ", err)
		return
	}
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPost, err))

	// Parsing result
	fmt.Println("Registered new user ", user, "and ID ", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
