package users

import (
	"encoding/json"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

// ChangeStatus a users: replaces ip and port by unspecified and 0
func ChangeStatus(w http.ResponseWriter, r *http.Request) {

	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var statusChange database.StatusChange
	err := decoder.Decode(&statusChange)
	if err != nil {
		log.Println("Error when decoding status change: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch statusChange.Status {
	case database.Online:
		err = database.SetStatusOnline(statusChange.ID)
		break
	case database.Offline:
		err = database.SetStatusOffline(statusChange.ID)
		break
	case database.Away:
		err = database.SetStatusAway(statusChange.ID)
		break
	case database.Busy:
		err = database.SetStatusBusy(statusChange.ID)
		break
	default:
		log.Println("Unrecognized status: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// SQL Queries
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPatch, err))
}
