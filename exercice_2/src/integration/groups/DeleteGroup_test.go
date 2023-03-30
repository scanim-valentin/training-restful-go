package groups

// Basic imports
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/groups"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type DeleteGroupTestSuite struct {
	suite.Suite
	usersInContacts []database.UserID
	method          string
	path            string
	nbGroups        int
	groupID         database.GroupID
	nbCharName      int
}

func (suite *DeleteGroupTestSuite) SetupTest() {

	suite.nbGroups = 100
	suite.groupID = 1
	suite.nbCharName = 50

	suite.method = "REMOVE"
	suite.path = "groups"
	database.ConnectDB("../../database/config_test.json")
	suite.usersInContacts = make([]database.UserID, 0)
	for k := 1; k < suite.nbGroups+1; k++ {
		group := database.RandomGroup(database.GroupID(k), suite.nbCharName)
		database.NewGroup(group)
	}
}

func (suite *DeleteGroupTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request
	for _, groupID := range suite.usersInContacts {
		req, err = http.NewRequest(suite.method,
			suite.path+"?id="+fmt.Sprintf("%v", groupID), nil)
		if err != nil {
			t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(groups.DeleteGroup)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func (suite *DeleteGroupTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request
	req, err = http.NewRequest(suite.method,
		suite.path+"?bad", nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(groups.DeleteGroup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *DeleteGroupTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeleteGroupTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteGroupTestSuite))
}
