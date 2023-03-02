package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"
	"service/utils"
)

// Register Registers a user with a username and save ip and port
func Register(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Registering new user with name", values["name"][0])
	ip, port := utils.GetIP(r)
	fmt.Println("Extracted IP from request ", ip, port)

	// SQL Queries
	id := database.InsertNewUser(values["name"][0], ip, port)

	// Parsing result
	fmt.Println("Registered new user with name ", values["name"][0], "and ID ", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.LoginResponse{ID: database.UserID(id), Username: values["name"][0], UserList: database.GetUserList()})
}
