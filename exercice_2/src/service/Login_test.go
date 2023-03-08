package service

// Basic imports
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"service/database"
	"testing"

	"github.com/stretchr/testify/suite"
)

const nbLoginUser = 100

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type LoginTestSuite struct {
	suite.Suite
	loginResponse database.LoginResponse
}

func (suite *LoginTestSuite) SetupTest() {
	database.ConnectDB("../database/config_test.json")
	suite.loginResponse.UserList = make([]database.User, 0)
	// Generating users
	for k := 1; k < nbLoginUser; k++ {
		newUser := database.RandomUser(database.UserID(k), 50)
		suite.loginResponse.UserList = append(suite.loginResponse.UserList, newUser)

		// Populating
		database.NewUser(newUser.Name, newUser.IP, fmt.Sprint(newUser.Port))
	}

	suite.loginResponse.ID = database.UserID(rand.Intn(nbLoginUser) + 1)
	suite.loginResponse.Username = suite.loginResponse.UserList[suite.loginResponse.ID-1].Name
}

// TestOk checks if signing in as an existing user provides a correct database.LoginResponse struct */
func (suite *LoginTestSuite) TestOk() {
	t := suite.T()
	url := fmt.Sprintf("/login?id=%v", suite.loginResponse.ID)
	req, err := http.NewRequest("GET", url, nil)
	addr := fmt.Sprintf("[%v]:%v", suite.loginResponse.UserList[suite.loginResponse.ID-1].IP, suite.loginResponse.UserList[suite.loginResponse.ID-1].Port)
	req.RemoteAddr = addr
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
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

// TestNotFound checks if signing in as an unknown user causes the expected HTTP error */

func (suite *LoginTestSuite) TestNotFound() {
	t := suite.T()
	url := fmt.Sprintf("/login?id=%v", nbLoginUser+1)
	req, err := http.NewRequest("GET", url, nil)
	addr := fmt.Sprintf("[%v]:%v", suite.loginResponse.UserList[suite.loginResponse.ID-1].IP, suite.loginResponse.UserList[suite.loginResponse.ID-1].Port)
	req.RemoteAddr = addr
	if err != nil {
		t.Fatal("Failed to initiate new http request: ", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func (suite *LoginTestSuite) TearDownTest() {
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
func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
