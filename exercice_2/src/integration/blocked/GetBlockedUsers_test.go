package blocked

// Basic imports
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/blocked"
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
type GetBlockedUsersTestSuite struct {
	suite.Suite
	BlockedUsers []database.User
	method       string
	path         string
}

func (suite *GetBlockedUsersTestSuite) SetupTest() {
	suite.method = "GET"
	suite.path = "blocked"
	database.ConnectDB("../../database/config_test.json")
	suite.BlockedUsers = make([]database.User, 0)
	for k := 1; k < nbUsers+1; k++ {
		user := database.RandomUser(database.UserID(k), nbCharName)
		user.Status = database.Offline
		database.NewUser(user.Name)
		if rand.Intn(proportionRange) == 0 {
			database.BlockUser(database.Block{UserID: userID, BlockedID: user.ID})
			suite.BlockedUsers = append(suite.BlockedUsers, user)
		}
	}
}

func (suite *GetBlockedUsersTestSuite) TestStatusOK() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?id="+fmt.Sprintf("%v", userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blocked.GetBlockedUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Extracting Data
	decoder := json.NewDecoder(rr.Body)
	var usersInContacts []database.User
	err = decoder.Decode(&usersInContacts)
	fmt.Println("Blocked users: ", usersInContacts)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(usersInContacts, suite.BlockedUsers) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			usersInContacts, suite.BlockedUsers)
	}

}

func (suite *GetBlockedUsersTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?nulnul="+fmt.Sprintf("%v", userID), nil)
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" contacts request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blocked.GetBlockedUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *GetBlockedUsersTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetBlockedUsersTestSuite(t *testing.T) {
	suite.Run(t, new(GetBlockedUsersTestSuite))
}
