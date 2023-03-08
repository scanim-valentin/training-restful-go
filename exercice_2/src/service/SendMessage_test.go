package service

// Basic imports
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service/database"
	"testing"

	"github.com/stretchr/testify/suite"
)

const messageContentLength = 50
const user1 = database.UserID(1)
const user2 = database.UserID(2)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SendMessageTestSuite struct {
	suite.Suite
}

func (suite *SendMessageTestSuite) SetupTest() {
	database.ConnectDB("../database/config_test.json")
}

// TestRegister checks if signing up as an existing user provides a correct database.LoginResponse struct */
func (suite *SendMessageTestSuite) TestCreated() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.RandomMessage(database.MessageID(0), user1, user2, messageContentLength))
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", "/send", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func (suite *SendMessageTestSuite) TestInternalStatusError() {
	if _, err := database.DB.Exec("DROP TABLE messages ; "); err != nil {
		suite.T().Fatal("Failed to drop table messages: ", err)
	}
	if _, err := database.DB.Exec("DROP TABLE users ; "); err != nil {
		suite.T().Fatal("Failed to drop table messages: ", err)
	}
	database.Close()
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.RandomMessage(database.MessageID(0), user1, user2, messageContentLength))
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", "/send", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSendMessageTestSuite(t *testing.T) {
	suite.Run(t, new(SendMessageTestSuite))
}
