package members

// Basic imports
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/groups/members"
	"service/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type GetUsersInGroupTestSuite struct {
	suite.Suite
	users         []database.User
	method        string
	path          string
	nbUsers       int
	groupID       database.GroupID
	nbCharName    int
	handler       func(w http.ResponseWriter, r *http.Request)
	ratio         int
	maxNameLength int
}

func (suite *GetUsersInGroupTestSuite) SetupTest() {

	suite.nbUsers = 500
	suite.groupID = 1
	suite.nbCharName = 50
	suite.ratio = 10 // ~50 users
	suite.handler = members.GetUsersInGroup

	suite.method = http.MethodGet
	suite.path = "groups/members"
	database.ConnectDB("../../../database/config_test.json")
	suite.users = make([]database.User, 0)
	suite.maxNameLength = 50
	suite.groupID = 0
	var err error
	if suite.groupID, err = database.NewGroup(database.Group{ID: suite.groupID, Name: string(utils.RandomString(suite.maxNameLength))}); err != nil {
		log.Fatal("Couldn't create group: ", err)
	}
	for k := 1; k < suite.nbUsers+1; k++ {
		user := database.RandomUser(database.UserID(k), suite.nbCharName)
		if user.ID, err = database.NewUser(user); err != nil {
			log.Fatal("Couldn't create user: ", err)
		}
		if err = database.SetStatus(user.ID, user.Status); err != nil {
			log.Fatal("Couldn't set status: ", err)
		}

		if rand.Intn(suite.ratio) == 0 {
			database.AddUserToGroup(database.Member{UserID: user.ID, GroupID: suite.groupID})
			suite.users = append(suite.users, user)
		}

	}
}

func (suite *GetUsersInGroupTestSuite) TestStatusOK() {
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
	var users []database.User
	err = decoder.Decode(&users)
	fmt.Println("Users in group: ", users)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(users, suite.users) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			users, suite.users)
	}

}

func (suite *GetUsersInGroupTestSuite) TestStatusBadRequest() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?bad="+fmt.Sprintf("%v", suite.groupID), nil)
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

func (suite *GetUsersInGroupTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetUsersInGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GetUsersInGroupTestSuite))
}
