package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestAddUserToContact(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock

	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}
	// TODO

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}

func TestRemoveUserFromContact(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock

	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}
	// TODO

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}
