package members

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func RemoveFromGroup(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var userid database.UserID
	value, ok := values["userid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &userid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan parameter: ", err)
		return
	}

	var groupid database.GroupID
	value, ok = values["groupid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &groupid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to scan parameter: ", err)
		return
	}

	// Call to database
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodDelete, database.RemoveUserFromGroup(database.Member{UserID: userid, GroupID: groupid})))
}
