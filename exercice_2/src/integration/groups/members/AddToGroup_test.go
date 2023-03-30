package members

// Basic imports
import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/groups/members"
	"service/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AddToGroupTestSuite struct {
	suite.Suite
	method        string
	path          string
	handler       func(w http.ResponseWriter, r *http.Request)
	groupID       database.GroupID
	userID        database.UserID
	maxNameLength int
}

func (suite *AddToGroupTestSuite) SetupTest() {
	suite.method = http.MethodPost
	suite.path = "groups/members"
	suite.handler = members.AddToGroup
	database.ConnectDB("../../../database/config_test.json")
	suite.maxNameLength = 50
	var err error
	suite.groupID = 0
	if suite.groupID, err = database.NewGroup(database.Group{ID: suite.groupID, Name: string(utils.RandomString(suite.maxNameLength))}); err != nil {
		log.Fatal("Couldn't register new group ", err)
	}
	if suite.userID, err = database.NewUser(database.RandomUser(suite.userID, suite.maxNameLength)); err != nil {
		log.Fatal("Couldn't register new user ", err)
	}
}

// TestOk checks if signing in as an existing users provides a correct database.LoginResponse struct */
func (suite *AddToGroupTestSuite) TestStatusCreated() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.Member{UserID: suite.userID, GroupID: suite.groupID})
	if err != nil {
		t.Fatal("Error when marshaling random member: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal("Fatal error when reading"+suite.method+"usersInContacts request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func (suite *AddToGroupTestSuite) TestStatusBadRequest() {
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
		t.Fatal("Fatal error when reading"+suite.method+"usersInContacts request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *AddToGroupTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAddToGroupTestSuite(t *testing.T) {
	suite.Run(t, new(AddToGroupTestSuite))
}
