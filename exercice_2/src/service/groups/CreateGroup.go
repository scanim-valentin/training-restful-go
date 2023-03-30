package groups

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var group database.Group
	err := decoder.Decode(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to read group in body request: ", err)
		return
	}
	var id database.GroupID
	id, err = database.NewGroup(group)
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPost, err))

	// Parsing result
	fmt.Println("Registered new group with name ", group.Name, "and ID ", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.Group{ID: id, Name: group.Name})
}
