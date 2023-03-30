package contacts

// Basic imports
import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/contacts"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type DeleteContactsTestSuite struct {
	suite.Suite
	usersInContacts []database.UserID
	method          string
	path            string
}

func (suite *DeleteContactsTestSuite) SetupTest() {
	suite.method = "REMOVE"
	suite.path = "blocked"
	database.ConnectDB("../../database/config_test.json")
	suite.usersInContacts = make([]database.UserID, 0)
	for k := 1; k < nbUsers+1; k++ {
		user := database.RandomUser(database.UserID(k), nbCharName)
		user.Status = database.Online
		database.NewUser(user)
		if rand.Intn(proportionRange) == 0 {
			database.AddUserToContacts(database.Contact{UserID: userID, ContactID: user.ID})
			suite.usersInContacts = append(suite.usersInContacts, user.ID)
		}
	}
}

func (suite *DeleteContactsTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request
	for _, contactID := range suite.usersInContacts {
		req, err = http.NewRequest(suite.method,
			suite.path+"?userid="+fmt.Sprintf("%v", userID)+"&contactid="+fmt.Sprintf("%v", contactID), nil)
		if err != nil {
			t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(contacts.RemoveFromContacts)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func (suite *DeleteContactsTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request
	req, err = http.NewRequest(suite.method,
		suite.path+"?bad", nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.RemoveFromContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *DeleteContactsTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeleteContactsTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteContactsTestSuite))
}
