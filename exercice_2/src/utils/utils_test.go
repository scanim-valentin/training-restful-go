package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"encoding/json"
)

func TestSend(t *testing.T) {

	// Setting up database
	Setup()

	// Formating JSON
	message := Message{MessageID(1), UserID(1), UserID(2), MessageContent("HELLO WORLD"), time.Now()}

	jsonenc := new(bytes.Buffer)
	json.NewEncoder(jsonenc).Encode(message)

	// Sending request
	req, err := http.NewRequest("POST", "/send", jsonenc)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(rr, req)

	// Testing
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
