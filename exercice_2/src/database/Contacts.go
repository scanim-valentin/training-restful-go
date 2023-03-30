package database

import (
	"database/sql"
	"log"
	"math/rand"
)

type Contact struct {
	UserID    UserID
	ContactID UserID
}

func AddUserToContacts(contact Contact) error {
	userID := contact.UserID
	contactID := contact.ContactID
	_, err := DB.Exec("INSERT INTO contacts (UserID, ContactID) VALUES ($1, $2)", userID, contactID)
	if err != nil {
		log.Println("Error when trying to insert new user to contacts: ", err)
	}
	return err
}

func RemoveUserFromContacts(userID UserID, contactID UserID) error {
	_, err := DB.Exec("DELETE FROM contacts WHERE UserID=$1 AND ContactID=$2", userID, contactID)
	return err
}

// GetContacts Retrieve online users list
func GetContacts(userID UserID) ([]User, error) {
	// SQL Queries
	// Order must be specified to match test specification
	// There is no default order (see https://stackoverflow.com/questions/6585574/postgres-default-sort-by-id-worldship "Rows are returned in an unspecified order")
	rows, err := DB.Query("SELECT id, username, status FROM users INNER JOIN contacts ON UserID=$1 AND ContactID=id ORDER BY id", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}(rows)

	var contacts []User

	// Reading rows
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Status); err != nil {
			break
		}
		contacts = append(contacts, user)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error in returned rows: ", err)
	}
	return contacts, err
}

func RandomContact(nbMaxUser int) Contact {
	return Contact{
		UserID:    UserID(rand.Intn(nbMaxUser - 1)),
		ContactID: UserID(rand.Intn(nbMaxUser - 1)),
	}
}
