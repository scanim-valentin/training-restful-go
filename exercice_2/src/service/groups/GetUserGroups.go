package groups

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func GetUserGroups(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var id database.UserID

	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan id parameter: ", err)
		return
	}

	groups, err := database.GetUserGroups(id)
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodGet, err))

	// Parsing result
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}
}
