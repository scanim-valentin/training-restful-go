package users

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var userID database.UserID
	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &userID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan parameter: ", err)
		return
	}
	err := database.DeleteUser(userID)
	// Call to database
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodDelete, err))
}
