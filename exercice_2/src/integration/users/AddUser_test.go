package users

// Basic imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"service/service/users"
	"service/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AddUserTestSuite struct {
	suite.Suite
	handler     func(w http.ResponseWriter, r *http.Request)
	method      string
	path        string
	user        database.User
	nbUsers     int
	maxCharName int
}

func (suite *AddUserTestSuite) SetupTest() {
	database.ConnectDB("../../database/config_test.json")
	suite.handler = users.AddUser
	suite.method = http.MethodPost
	suite.path = "/users"
	suite.nbUsers = 50
	suite.maxCharName = 50
	suite.user.ID = database.UserID(suite.nbUsers + 1)
	suite.user = database.RandomUser(suite.user.ID, suite.maxCharName)
	// Generating users
	for k := 1; k < suite.nbUsers+1; k++ {
		database.NewUser(database.RandomUser(database.UserID(k), 50))
	}
}

// TestRegister checks if signing up as an existing users provides a correct database.LoginResponse struct */
func (suite *AddUserTestSuite) TestOk() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(suite)
	if err != nil {
		t.Fatal("Error when marshaling random contact: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got \n%v want \n%v",
			status, http.StatusCreated)
	}
	// Check the response body is what we expect.

	decoder := json.NewDecoder(rr.Body)
	var userID database.UserID
	err = decoder.Decode(&userID)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(userID, suite.user.ID) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			userID, suite.user.ID)
	}

}

func (suite *AddUserTestSuite) TestBadRequest() {
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
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got \n%v want \n%v",
			status, http.StatusBadRequest)
	}

}

func (suite *AddUserTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAddUserSuite(t *testing.T) {
	suite.Run(t, new(AddUserTestSuite))
}
