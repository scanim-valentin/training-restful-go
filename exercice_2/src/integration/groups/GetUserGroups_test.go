package groups

// Basic imports
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/groups"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type GetUserGroupsTestSuite struct {
	suite.Suite
	groups     []database.Group
	method     string
	path       string
	nbGroups   int
	userID     database.UserID
	nbCharName int
	handler    func(w http.ResponseWriter, r *http.Request)
}

func (suite *GetUserGroupsTestSuite) SetupTest() {

	suite.nbGroups = 100
	suite.userID = 1
	suite.nbCharName = 50
	suite.handler = groups.GetUserGroups

	suite.method = "GET"
	suite.path = "groups"
	database.ConnectDB("../../database/config_test.json")
	suite.groups = make([]database.Group, 0)
	for k := 1; k < suite.nbGroups+1; k++ {
		group := database.RandomGroup(database.GroupID(k), suite.nbCharName)
		database.AddUserToGroup(database.Member{UserID: suite.userID, GroupID: database.GroupID(k)})
		database.NewGroup(group)
		suite.groups = append(suite.groups, group)
	}
}

func (suite *GetUserGroupsTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?id="+fmt.Sprintf("%v", suite.userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
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
	var groups []database.Group
	err = decoder.Decode(&groups)
	fmt.Println("Users in contact: ", groups)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(groups, suite.groups) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			groups, suite.groups)
	}

}

func (suite *GetUserGroupsTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?nulnul="+fmt.Sprintf("%v", suite.userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *GetUserGroupsTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetUserGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserGroupsTestSuite))
}
