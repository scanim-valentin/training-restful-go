package messages

// Basic imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/contacts"
	"service/service/groups"
	"service/service/messages"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SendMessageTestSuite struct {
	suite.Suite
	method      string
	path        string
	nbMessages  int
	messageID   database.MessageID
	userID      database.UserID
	groupID     database.GroupID
	maxNameChar int
	handler     func(w http.ResponseWriter, r *http.Request)
}

func (suite *SendMessageTestSuite) SetupTest() {

	suite.nbMessages = 500
	suite.userID = 1
	suite.groupID = 1
	suite.maxNameChar = 50
	suite.handler = messages.SendMessage
	suite.method = http.MethodPost
	suite.path = "messages"
	database.ConnectDB("../../database/config_test.json")
	if _, err := database.NewGroup(database.RandomGroup(suite.groupID, suite.maxNameChar)); err != nil {
		log.Fatal("Couldn't create group: ", err)
	}
	var err error
	if suite.userID, err = database.NewUser(database.RandomUser(suite.userID, suite.maxNameChar)); err != nil {
		log.Fatal("Couldn't create userID: ", err)
	}
	if database.AddUserToGroup(database.Member{UserID: suite.userID, GroupID: suite.groupID}) != nil {
		log.Fatal("Couldn't create group: ", err)
	}
}

func (suite *SendMessageTestSuite) TestStatusOK() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.RandomMessage(suite.messageID, suite.userID, suite.groupID, suite.maxNameChar))
	if err != nil {
		t.Fatal("Error when marshaling random contact: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" messageID request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(groups.CreateGroup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	// Extracting Data
	decoder := json.NewDecoder(rr.Body)
	var messageID database.MessageID
	err = decoder.Decode(&messageID)
	fmt.Println("Users in contact: ", messageID)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(messageID, suite.messageID) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			messageID, suite.messageID)
	}
}

func (suite *SendMessageTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(utils.DummyStruct{Dummy: 1})
	if err != nil {
		t.Fatal("Error when marshaling random contact: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" groups request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.AddToContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *SendMessageTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSendMessageTestSuite(t *testing.T) {
	suite.Run(t, new(GetConversationTestSuite))
}
