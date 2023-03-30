package members

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/groups/members"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type RemoveFromGroupTestSuite struct {
	suite.Suite
	method        string
	path          string
	handler       func(w http.ResponseWriter, r *http.Request)
	groupID       database.GroupID
	userID        database.UserID
	maxNameLength int
}

func (suite *RemoveFromGroupTestSuite) SetupTest() {
	suite.method = http.MethodPost
	suite.path = "groups/members"
	suite.handler = members.RemoveFromGroup
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
	if err = database.AddUserToGroup(database.Member{UserID: suite.userID, GroupID: suite.groupID}); err != nil {
		log.Fatal("Couldn't add user to group ", err)
	}
}

// TestOk checks if signing in as an existing users provides a correct database.LoginResponse struct */
func (suite *RemoveFromGroupTestSuite) TestStatusCreated() {
	t := suite.T()
	var err error
	var req *http.Request
	req, err = http.NewRequest(suite.method, suite.path+"?userid="+fmt.Sprintf("%v", suite.userID)+"&groupid="+fmt.Sprintf("%v", suite.groupID), nil)
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

func (suite *RemoveFromGroupTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request
	req, err = http.NewRequest(suite.method, suite.path+"?bad", nil)
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

func (suite *RemoveFromGroupTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRemoveFromGroupTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveFromGroupTestSuite))
}
