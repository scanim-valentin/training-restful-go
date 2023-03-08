package service

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
)

// Logout a user: replaces ip and port by unspecified and 0
func Logout(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	fmt.Println("Login out user with ID ", values["id"][0])
	// SQL Queries
	var id database.UserID
	if _, err := fmt.Sscanf(values["id"][0], "%d", &id); err != nil {
		log.Panic(err)
	}
	database.SetStatusOffline(id)
}
