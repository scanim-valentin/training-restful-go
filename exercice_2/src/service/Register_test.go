package service

// Basic imports
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"testing"

	"github.com/stretchr/testify/suite"
)

const nbRegisterUser = 100
const nameLength = 50

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type RegisterTestSuite struct {
	suite.Suite
	loginResponse database.LoginResponse
}

func (suite *RegisterTestSuite) SetupTest() {
	database.ConnectDB("../database/config_test.json")
	suite.loginResponse.UserList = make([]database.User, 0)
	// Generating users
	for k := 1; k < nbLoginUser+1; k++ {
		newUser := database.RandomUser(database.UserID(k), 50)
		suite.loginResponse.UserList = append(suite.loginResponse.UserList, newUser)

		// Populating
		if k < nbLoginUser {
			database.InsertNewUser(newUser.Name, newUser.IP, fmt.Sprint(newUser.Port))

		}
	}

	suite.loginResponse.ID = database.UserID(nbLoginUser)
	suite.loginResponse.Username = suite.loginResponse.UserList[suite.loginResponse.ID-1].Name
}

// TestRegister checks if signing up as an existing user provides a correct database.LoginResponse struct */
func (suite *RegisterTestSuite) TestOk() {
	t := suite.T()
	url := fmt.Sprintf("/register?name=%v", suite.loginResponse.Username)
	req, err := http.NewRequest("GET", url, nil)
	addr := fmt.Sprintf("[%v]:%v", suite.loginResponse.UserList[suite.loginResponse.ID-1].IP, suite.loginResponse.UserList[suite.loginResponse.ID-1].Port)
	req.RemoteAddr = addr
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.

	decoder := json.NewDecoder(rr.Body)
	var loginResponse database.LoginResponse
	err = decoder.Decode(&loginResponse)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	if !reflect.DeepEqual(loginResponse, suite.loginResponse) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			loginResponse, suite.loginResponse)
	}

}

func (suite *RegisterTestSuite) TearDownTest() {
	if _, err := database.DB.Exec("DROP TABLE messages ; "); err != nil {
		suite.T().Fatal("Failed to drop table messages: ", err)
	}
	if _, err := database.DB.Exec("DROP TABLE users ; "); err != nil {
		suite.T().Fatal("Failed to drop table messages: ", err)
	}
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
