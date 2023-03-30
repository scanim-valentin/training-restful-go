package blocked

// Basic imports
import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/blocked"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type UnblockUserTestSuite struct {
	suite.Suite
	usersInContacts []database.UserID
	method          string
	path            string
}

func (suite *UnblockUserTestSuite) SetupTest() {
	suite.method = "DELETE"
	suite.path = "blocked"
	database.ConnectDB("../../database/config_test.json")
	suite.usersInContacts = make([]database.UserID, 0)
	for k := 1; k < nbUsers+1; k++ {
		user := database.RandomUser(database.UserID(k), nbCharName)
		user.Status = database.Online
		database.NewUser(user.Name)
		if rand.Intn(proportionRange) == 0 {
			database.AddUserToContacts(database.Contact{UserID: userID, ContactID: user.ID})
			suite.usersInContacts = append(suite.usersInContacts, user.ID)
		}
	}
}

func (suite *UnblockUserTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request
	for _, blockedID := range suite.usersInContacts {
		req, err = http.NewRequest(suite.method, suite.path+"?userid="+fmt.Sprintf("%v", userID)+"&blockedid="+fmt.Sprintf("%v", blockedID), nil)
		if err != nil {
			t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(blocked.UnblockUser)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func (suite *UnblockUserTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request
	req, err = http.NewRequest(suite.method, suite.path+"?bad", nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blocked.UnblockUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *UnblockUserTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUnblockUserTestSuite(t *testing.T) {
	suite.Run(t, new(UnblockUserTestSuite))
}
