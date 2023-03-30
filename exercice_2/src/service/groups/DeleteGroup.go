package groups

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var groupID database.GroupID

	value, ok := values["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan id parameter: ", err)
		return
	}

	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodGet, database.DeleteGroup(groupID)))
}
