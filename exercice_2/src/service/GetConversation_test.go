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

const nbMessages = 500
const nbUser = 100
const nbMaxChar = 100

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	conversation database.Conversation
}

func (suite *ExampleTestSuite) SetupTest() {
	database.ConnectDB("../database/config_test.json")
	messageList := make([]database.Message, 0)
	conversation := make([]database.Message, 0)
	// fmt.Println("Generating random messages...")
	var randomMessage database.Message
	for k := 1; k < nbMessages+1; k++ {
		// Randomly deciding whether the message is part of the conversation between 1 and 2
		if rand.Intn(2) == 0 {
			source := 1
			destination := 2
			// Randomly deciding whether the message is from 1 to 2 or from 2 to 1
			if rand.Intn(2) == 0 {
				source = 2
				destination = 1
			}
			randomMessage = database.RandomMessage(database.MessageID(k), database.UserID(source), database.UserID(destination), nbMaxChar)
			conversation = append(conversation, randomMessage)
		} else {
			randomMessage = database.RandomMessage(database.MessageID(k), database.UserID(rand.Intn(nbUser-2)+2), database.UserID(rand.Intn(nbUser-2)+2), nbMaxChar)
		}
		// Populating
		if _, err := database.NewMessage(randomMessage.Source, randomMessage.Destination, randomMessage.Content, randomMessage.Time); err != nil {
			suite.T().Fatal("Failed to populate database: ", err)
		}
	}
	// fmt.Println("Successfully populated database with", nbMessages, "messages")
	messageList = append(messageList, randomMessage)
	suite.conversation = database.Conversation{
		Messages: conversation,
	}

}

// TestExample TODO: issue with time formating (Postgres precision). Evaluate time.Now exclusively in Message.NewMessage ?  */
func (suite *ExampleTestSuite) TestExample() {
	t := suite.T()
	req, err := http.NewRequest("GET", "/select?user=1&other=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetConversation)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		fmt.Println()
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	// Extracting Data
	decoder := json.NewDecoder(rr.Body)
	var conversation database.Conversation
	err = decoder.Decode(&conversation)
	if err != nil {
		t.Fatal("Unable to decode json reply: ", err)
	}

	fmt.Println("\nConversation: ", conversation)
	fmt.Println("\nConversation truth: ", suite.conversation)

	if !reflect.DeepEqual(conversation, suite.conversation) {
		fmt.Println()
		t.Errorf("handler returned unexpected body: \n got %v \n want %v",
			conversation, suite.conversation)
	}
}

func (suite *ExampleTestSuite) TearDownTest() {
	if _, err := database.DB.Exec("DROP TABLE messages ; "); err != nil {
		suite.T().Fatal("Failed to drop table messages: ", err)
	}
	database.Close()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
