package contacts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	var id database.UserID
	if _, err := fmt.Sscanf(value[0], "%d", &id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan id parameter: ", err)
		return
	}
	contacts, err := database.GetContacts(id)
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodGet, err))

	// Parsing result
	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
