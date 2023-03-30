package blocked

import (
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func UnblockUser(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	values := r.URL.Query()
	var userID, blockedID database.UserID

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

	value, ok = values["blockedid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Value id cannot be found in URL")
		return
	}
	if _, err := fmt.Sscanf(value[0], "%d", &blockedID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Failed to scan blockedID parameter: ", err)
	}

	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodDelete,
		database.UnblockUser(database.Block{UserID: database.UserID(userID), BlockedID: database.UserID(blockedID)})))
}
