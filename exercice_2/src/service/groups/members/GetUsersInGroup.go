package members

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func GetUsersInGroup(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var id database.GroupID
	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan parameter: ", err)
		return
	}

	users, err := database.GetUsersInGroup(id)

	// Parsing result
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Couldn't encode : ", err)
		return
	}

	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodGet, err))
}
