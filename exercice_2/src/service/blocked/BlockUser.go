package blocked

import (
	"encoding/json"
	"log"
	"net/http"
	"service/database"
	"service/utils"
)

func BlockUser(w http.ResponseWriter, r *http.Request) {
	// Extracting Data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var block database.Block
	err := decoder.Decode(&block)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to read blocked user in body request: ", err)
		return
	}

	// Call to database
	w.WriteHeader(utils.SQLErrorToHTTPStatus(http.MethodPost, database.BlockUser(block)))
}
