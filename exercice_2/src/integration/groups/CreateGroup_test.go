package groups

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/contacts"
	"service/service/groups"
	"service/utils"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type CreateGroupTestSuite struct {
	suite.Suite
	method   string
	path     string
	nbGroups int
	groupID  database.GroupID
}

func (suite *CreateGroupTestSuite) SetupTest() {

	suite.nbGroups = 100
	suite.groupID = 1

	suite.method = "POST"
	suite.path = "/groups"
	database.ConnectDB("../../database/config_test.json")
}

// TestOk checks if signing in as an existing users provides a correct database.LoginResponse struct */
func (suite *CreateGroupTestSuite) TestStatusCreated() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.RandomGroup(suite.groupID, suite.nbGroups))
	if err != nil {
		t.Fatal("Error when marshaling random contact: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal("Fatal error when reading "+suite.method+" groups request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(groups.CreateGroup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func (suite *CreateGroupTestSuite) TestStatusBadRequest() {
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
		t.Fatal("Fatal error when reading "+suite.method+" groups request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contacts.AddToContacts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *CreateGroupTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCreateGroupTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGroupTestSuite))
}
