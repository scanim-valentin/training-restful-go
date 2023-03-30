package contacts

// Basic imports
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/contacts"
	"testing"

	"github.com/stretchr/testify/suite"
)

const nbUsers int = 500
const nbCharName int = 20
const proportionRange int = 5
const userID database.UserID = 0

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type GetContactsTestSuite struct {
	suite.Suite
	usersInContacts []database.User
	method          string
	path            string
}

func (suite *GetContactsTestSuite) SetupTest() {
	suite.method = "GET"
	suite.path = "blocked"
	database.ConnectDB("../../database/config_test.json")
	suite.usersInContacts = make([]database.User, 0)
	for k := 1; k < nbUsers+1; k++ {
		user := database.RandomUser(database.UserID(k), nbCharName)
		user.Status = database.Online
		database.NewUser(user)
		if rand.Intn(proportionRange) == 0 {
			database.AddUserToContacts(database.Contact{UserID: userID, ContactID: user.ID})
			suite.usersInContacts = append(suite.usersInContacts, user)
		}
	}
}

func (suite *GetContactsTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?id="+fmt.Sprintf("%v", userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.GetContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Extracting Data
	decoder := json.NewDecoder(rr.Body)
	var usersInContacts []database.User
	err = decoder.Decode(&usersInContacts)
	fmt.Println("Users in contact: ", usersInContacts)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(usersInContacts, suite.usersInContacts) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			usersInContacts, suite.usersInContacts)
	}

}

func (suite *GetContactsTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?nulnul="+fmt.Sprintf("%v", userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.GetContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *GetContactsTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetContactsTestSuite(t *testing.T) {
	suite.Run(t, new(GetContactsTestSuite))
}
