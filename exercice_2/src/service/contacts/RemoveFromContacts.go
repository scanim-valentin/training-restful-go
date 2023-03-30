package contacts

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func RemoveFromContacts(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var userID, contactID database.UserID

	value, ok := values["userid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &userID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Failed to scan userid parameter: ", err)
	}

	value, ok = values["contactid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &contactID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Failed to scan contactID parameter: ", err)
	}

	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodDelete, database.RemoveUserFromContacts(userID, contactID)))
}
