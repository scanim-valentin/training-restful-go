package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"service/utils"
	"testing"
)

const maxNameLength = 50
const nbNewUsers = 500
const nbLoginUser = 500

func TestGetUserList(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock
	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "ip", "port"})

	// Generating users
	for k := 0; k < nbNewUsers; k++ {
		newUser := RandomUser(UserID(k), maxNameLength)
		mock.
			ExpectExec("INSERT INTO users (.+)").
			WithArgs(fmt.Sprint(newUser.Name), fmt.Sprint(utils.IPUnspecified), fmt.Sprint(newUser.Port))

		DB.Exec("INSERT INTO users (username, ip, port) VALUES ($1, $2, $3)", fmt.Sprint(newUser.Name), fmt.Sprint(utils.IPUnspecified), fmt.Sprint(newUser.Port))
		rows.AddRow(k, newUser.Name, utils.IPUnspecified, newUser.Port)
	}
	mock.
		ExpectQuery("SELECT (.+) FROM users ORDER BY id").
		WillReturnRows(rows)
	GetUserList()
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}

func TestLoginUser(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock

	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}

	// Generating users
	for k := 1; k < nbLoginUser; k++ {
		newUser := RandomUser(UserID(k), maxNameLength)

		mock.
			ExpectExec("INSERT INTO users (.+)").
			WithArgs(fmt.Sprint(newUser.Name), fmt.Sprint(utils.IPUnspecified), fmt.Sprint(newUser.Port))

		DB.Exec("INSERT INTO users (username, ip, port) VALUES ($1, $2, $3)", fmt.Sprint(newUser.Name), fmt.Sprint(utils.IPUnspecified), fmt.Sprint(newUser.Port))

		mock.
			ExpectQuery("UPDATE users SET (.+) RETURNING username").
			WithArgs(newUser.IP.String(), fmt.Sprint(newUser.Port), newUser.ID).
			WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow(newUser.Name))

		LoginUser(newUser.IP, fmt.Sprint(newUser.Port), newUser.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}

func TestNewUser(t *testing.T) {
	var err error
	var mock sqlmock.Sqlmock

	DB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal("Failed to initialise test database: ", err)
	}

	for k := 0; k < nbNewUsers; k++ {
		newUser := RandomUser(UserID(0), maxNameLength)
		mock.
			ExpectQuery("INSERT INTO users (.+) RETURNING id").
			WithArgs(fmt.Sprint(newUser.Name), fmt.Sprint(newUser.IP), fmt.Sprint(newUser.Port)).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(k))
		NewUser(newUser.Name, newUser.IP, fmt.Sprint(newUser.Port))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations were not met: ", err)
	}
}
