package database

import (
	"log"
	"service/utils"
)

// UserID Unique users identifier
type UserID int

// User A users has a unique UserID, a name (not unique), an IP address and a Port.
// IP and Port are also used to indicate the users's status (connected or not)
type User struct {
	ID     UserID
	Name   string
	Status string
}

// LoginResponse Login response is returned to login and register HTTP requests
type LoginResponse struct {
	ID       UserID
	Username string
	Contacts []User
}

type StatusChange struct {
	ID     UserID
	Status string
}

// NewUser inserts new users and returns a unique newly created UserID
func NewUser(user User) (UserID, error) {
	name := user.Name
	var id UserID
	err := DB.QueryRow("INSERT INTO users (username, status) VALUES ($1, $2) RETURNING id", name, Online).Scan(&id)
	if err != nil {
		log.Println("Error when inserting new user: ", err)
	}
	return id, err
}

func LoginUser(id UserID) (string, error) {
	var name string
	err := DB.QueryRow("UPDATE users SET status = $1 WHERE id = $2 RETURNING username", Online, id).Scan(&name)
	return name, err
}

func SetStatus(id UserID, status string) error {
	_, err := DB.Exec("UPDATE users SET status = $1 WHERE id = $2", status, id)
	return err
}

func SetStatusOnline(id UserID) error {
	return SetStatus(id, Online)
}

func SetStatusOffline(id UserID) error {
	return SetStatus(id, Offline)
}

func SetStatusBusy(id UserID) error {
	return SetStatus(id, Busy)
}

func SetStatusAway(id UserID) error {
	return SetStatus(id, Away)
}

func DeleteUser(id UserID) error {
	_, err := DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Couldn't delete user: ", err)
	}
	return err
}

func RandomUser(id UserID, nbMaxChar int) User {
	return User{
		ID:     id,
		Name:   string(utils.RandomString(nbMaxChar)),
		Status: RandomStatus(),
	}
}
