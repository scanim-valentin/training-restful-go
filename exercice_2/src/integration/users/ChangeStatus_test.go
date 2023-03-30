package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/users"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ChangeStatusTestSuite struct {
	suite.Suite
	handler     func(w http.ResponseWriter, r *http.Request)
	method      string
	path        string
	userID      database.UserID
	maxCharName int
}

func (suite *ChangeStatusTestSuite) SetupTest() {
	database.ConnectDB("../../database/config_test.json")
	suite.handler = users.ChangeStatus
	suite.method = http.MethodPatch
	suite.path = "/users"
	suite.maxCharName = 50
	suite.userID = 1
	database.NewUser(database.RandomUser(suite.userID, suite.maxCharName))
}

func GenericTestStatus(suite *ChangeStatusTestSuite, status string) {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.StatusChange{ID: suite.userID, Status: status})
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
	if status := rr.Code; status != http.StatusOK {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got \n%v want \n%v",
			status, http.StatusOK)
	}

}

func (suite *ChangeStatusTestSuite) TestOnline() {
	GenericTestStatus(suite, database.Online)
}
func (suite *ChangeStatusTestSuite) TestOffline() {
	GenericTestStatus(suite, database.Offline)
}
func (suite *ChangeStatusTestSuite) TestBusy() {
	GenericTestStatus(suite, database.Busy)
}
func (suite *ChangeStatusTestSuite) TestAway() {
	GenericTestStatus(suite, database.Away)
}

func (suite *ChangeStatusTestSuite) TestBadRequest() {
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

func (suite *ChangeStatusTestSuite) DeleteUserTestSuite() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestChangeStatusTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeStatusTestSuite))
}
