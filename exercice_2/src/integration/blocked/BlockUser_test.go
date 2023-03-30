package blocked

// Basic imports
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service/database"
	"service/service/blocked"
	"service/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

const nbContacts = 100

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BlockTestSuite struct {
	suite.Suite
	method string
	path   string
}

func (suite *BlockTestSuite) SetupTest() {
	suite.method = "POST"
	suite.path = "blocked"
	database.ConnectDB(database.TestDBPath)
}

// TestOk checks if signing in as an existing users provides a correct database.LoginResponse struct */
func (suite *BlockTestSuite) TestStatusCreated() {
	t := suite.T()
	var jsonByte []byte
	var err error
	var req *http.Request
	jsonByte, err = json.Marshal(database.RandomBlock(nbContacts))
	if err != nil {
		t.Fatal("Error when marshaling random contact: ", err)
	}

	req, err = http.NewRequest(suite.method, suite.path, bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Fatal("Fatal error when reading POST BlockedUsers request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blocked.BlockUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func (suite *BlockTestSuite) TestStatusBadRequest() {
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
		t.Fatal("Fatal error when reading POST BlockedUsers request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blocked.BlockUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func (suite *BlockTestSuite) TearDownTest() {
	database.DropAllTables()
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAddUserToContactsTestSuite(t *testing.T) {
	suite.Run(t, new(BlockTestSuite))
}
