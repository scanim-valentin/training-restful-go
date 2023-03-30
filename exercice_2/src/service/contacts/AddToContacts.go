package contacts

import (
	"encoding/json"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

// AddToContacts adds a users with ID to
func AddToContacts(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var contact database.Contact
	err := decoder.Decode(&contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to read contact in body request: ", err)
		return
	} else {
		w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPost, database.AddUserToContacts(contact)))
	}
}
