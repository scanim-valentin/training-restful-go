package messages

// Basic imports
import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/messages"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type GetConversationTestSuite struct {
	suite.Suite
	messages    []database.Message
	method      string
	path        string
	nbMessages  int
	nbUsers     int
	nbGroups    int
	userID      database.UserID
	groupID     database.GroupID
	maxNameChar int
	ratio       int
	handler     func(w http.ResponseWriter, r *http.Request)
}

func (suite *GetConversationTestSuite) SetupTest() {

	suite.nbMessages = 500
	suite.nbUsers = 500
	suite.nbGroups = 500

	suite.userID = 1
	suite.groupID = 1
	suite.maxNameChar = 50
	suite.handler = messages.GetConversation
	suite.ratio = 10

	suite.method = "GET"
	suite.path = "messages"
	database.ConnectDB("../../database/config_test.json")
	suite.messages = make([]database.Message, 0)
	for k := 1; k < suite.nbGroups; k++ {
		if _, err := database.NewGroup(database.Group{ID: database.GroupID(k), Name: string(utils.RandomString(suite.maxNameChar))}); err != nil {
			log.Fatal("Couldn't create group: ", err)
		}
	}
	var err error
	if suite.userID, err = database.NewUser(database.RandomUser(suite.userID, suite.maxNameChar)); err != nil {
		log.Fatal("Couldn't create userID: ", err)
	}
	if database.AddUserToGroup(database.Member{UserID: suite.userID, GroupID: suite.groupID}) != nil {
		log.Fatal("Couldn't create group: ", err)
	}
	for k := 2; k < suite.nbUsers; k++ {
		var userID database.UserID
		var err error
		if userID, err = database.NewUser(database.RandomUser(database.UserID(k), suite.maxNameChar)); err != nil {
			log.Fatal("Couldn't create userID: ", err)
		}
		if database.AddUserToGroup(database.Member{UserID: userID, GroupID: database.GroupID(rand.Intn(suite.nbUsers-2) + 1)}) != nil {
			log.Fatal("Couldn't create group: ", err)
		}
	}
	for k := 1; k < suite.nbMessages+1; k++ {
		var message database.Message
		if rand.Intn(suite.ratio) == 0 {
			message = database.RandomMessage(database.MessageID(k), suite.userID, suite.groupID, suite.maxNameChar)
		} else {
			message = database.RandomMessage(database.MessageID(k), database.UserID(rand.Intn(suite.nbUsers-1)+1), database.GroupID(rand.Intn(suite.nbUsers-1)+1), suite.maxNameChar)
		}
		if message.ID, err = database.NewMessage(message); err != nil {
			log.Fatal("Couldn't create userID: ", err)
		}
		if message.User == suite.userID && message.Group == suite.groupID {
			fmt.Println("OK")
			suite.messages = append(suite.messages, message)
		}
	}
}

func (suite *GetConversationTestSuite) TestStatusOK() {
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
	// Extracting Data
	decoder := json.NewDecoder(rr.Body)
	var messages []database.Message
	err = decoder.Decode(&messages)
	fmt.Println("Users in contact: ", messages)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(messages, suite.messages) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			messages, suite.messages)
	}

}

func (suite *GetConversationTestSuite) TestStatusBadRequest() {
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

func (suite *GetConversationTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetConversationTestSuite(t *testing.T) {
	suite.Run(t, new(GetConversationTestSuite))
}
