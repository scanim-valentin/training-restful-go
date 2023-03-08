package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

// Login a user with a username and save ip and port
func Login(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login in user with ID ", values["id"][0])
	ip, port := utils.GetIP(r)
	fmt.Println("Extracted IP from request ", ip, port)
	// SQL Queries
	var id database.UserID
	if _, err := fmt.Sscanf(values["id"][0], "%d", &id); err != nil {
		log.Panic(err)
	}
	if namePtr := database.LoginUser(ip, port, id); namePtr != nil {
		// Parsing result
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(database.LoginResponse{ID: id, Username: string(*namePtr), UserList: database.GetUserList()})
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
