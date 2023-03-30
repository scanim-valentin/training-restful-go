package messages

// Basic imports
import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/messages"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type DeleteMessageSuite struct {
	suite.Suite
	message     database.Message
	method      string
	path        string
	userID      database.UserID
	groupID     database.GroupID
	maxNameChar int
	handler     func(w http.ResponseWriter, r *http.Request)
}

func (suite *DeleteMessageSuite) SetupTest() {

	suite.userID = 1
	suite.groupID = 1
	suite.maxNameChar = 50
	suite.handler = messages.DeleteMessage
	suite.message.ID = 1
	suite.method = "DELETE"
	suite.path = "messages"
	database.ConnectDB("../../database/config_test.json")
	if _, err := database.NewGroup(database.Group{ID: suite.groupID, Name: string(utils.RandomString(suite.maxNameChar))}); err != nil {
		log.Fatal("Couldn't create group: ", err)
	}
	var err error
	if suite.userID, err = database.NewUser(database.RandomUser(suite.userID, suite.maxNameChar)); err != nil {
		log.Fatal("Couldn't create userID: ", err)
	}
	if database.AddUserToGroup(database.Member{UserID: suite.userID, GroupID: suite.groupID}) != nil {
		log.Fatal("Couldn't create group: ", err)
	}
	suite.message = database.RandomMessage(suite.message.ID, suite.userID, suite.groupID, suite.maxNameChar)
	database.NewMessage(suite.message)
}

func (suite *DeleteMessageSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?id="+fmt.Sprintf("%v", suite.groupID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func (suite *DeleteMessageSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?nul", nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func (suite *DeleteMessageSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeleteMessageSuite(t *testing.T) {
	suite.Run(t, new(DeleteMessageSuite))
}
