package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
)

const nbMessages = 500
const nbMaxCharMessage = 100
const nbUser = 5000

func TestGetMessages(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock
	rows := sqlmock.NewRows([]string{"id", "source", "destination", "content", "time"})
	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Error while initializing new mock DB: ", err)
	}
	var randomMessage Message
	for k := 0; k < nbMessages; k++ {
		// Randomly deciding whether the message is part of the conversation between 1 and 2
		if rand.Intn(2) == 0 {
			user := 1
			group := 2
			// Randomly deciding whether the message is from 1 to 2 or from 2 to 1
			if rand.Intn(2) == 0 {
				user = 2
				group = 1
			}
			randomMessage = RandomMessage(MessageID(k), UserID(user), GroupID(group), nbMaxCharMessage)
			// Just making sure messages are sent at human-like rate in to preserve the time based order of a conversation
			randomMessage.Time = randomMessage.Time.Add(time.Duration(k) * time.Second)

			rows.AddRow(k, randomMessage.User, randomMessage.Group, randomMessage.Content, randomMessage.Time)

		} else {
			randomMessage = RandomMessage(MessageID(k), UserID(rand.Intn(nbUser-2)+2), GroupID(rand.Intn(nbUser-2)+2), nbMaxCharMessage)
		}
		// Populating
		mock.
			ExpectQuery("INSERT INTO messages (.+)").
			WithArgs(randomMessage.User, randomMessage.Group, randomMessage.Content, randomMessage.Time)

		DB.QueryRow("INSERT INTO messages (source, destination, content, time) VALUES ($1, $2, $3, $4)",
			randomMessage.User, randomMessage.Group, randomMessage.Content, randomMessage.Time)

	}
	mock.
		ExpectQuery("SELECT (.+) FROM messages WHERE (.+) ORDER BY time ASC").
		WithArgs(1, 2).
		WillReturnRows(rows).
		RowsWillBeClosed()
	GetMessages(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}

}

func TestNewMessage(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock

	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}

	for k := 0; k < nbNewUsers; k++ {
		randomMessage := RandomMessage(MessageID(0), UserID(0), GroupID(0), nbMaxCharMessage)
		mock.
			ExpectQuery("INSERT INTO messages (.+) RETURNING id").
			WithArgs(randomMessage.User, randomMessage.Group, randomMessage.Content, randomMessage.Time).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(k))
		id, err := NewMessage(randomMessage.User, randomMessage.Group, randomMessage.Content, randomMessage.Time)
		fmt.Println("New message ID: ", id)
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}
