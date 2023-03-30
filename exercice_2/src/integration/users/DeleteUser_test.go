package users

// Basic imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/users"
	"service/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type DeleteUserTestSuite struct {
	suite.Suite
	handler     func(w http.ResponseWriter, r *http.Request)
	method      string
	path        string
	userID      database.UserID
	maxCharName int
}

func (suite *DeleteUserTestSuite) SetupTest() {
	database.ConnectDB("../../database/config_test.json")
	suite.handler = users.DeleteUser
	suite.method = http.MethodDelete
	suite.path = "/users"
	suite.maxCharName = 50
	suite.userID = 1
	database.NewUser(database.RandomUser(suite.userID, suite.maxCharName))
}

// TestRegister checks if signing up as an existing users provides a correct database.LoginResponse struct */
func (suite *DeleteUserTestSuite) TestOk() {
	t := suite.T()
	var err error
	var req *http.Request

	req, err = http.NewRequest(suite.method, suite.path+"?id="+fmt.Sprint(suite.userID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got \n%v want \n%v",
			status, http.StatusOK)
	}

}

func (suite *DeleteUserTestSuite) TestBadRequest() {
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

func (suite *DeleteUserTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeleteUserTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserTestSuite))
}