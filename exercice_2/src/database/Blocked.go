package database

import (
	"database/sql"
	"log"
	"math/rand"
)

type Block struct {
	UserID    UserID
	BlockedID UserID
}

func GetBlockedUsers(id UserID) ([]User, error) {
	rows, err := DB.Query("SELECT id, username FROM users INNER JOIN blocked ON UserID=$1 AND BlockedID=id ORDER BY id", id)
	if err != nil {
		log.Panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic("Failed to close rows in GetMessage: ", err)
		}
	}(rows)
	blockedUsers := make([]User, 0)

	// Reading rows
	for rows.Next() {
		var user User
		// Status details should only be available for non-blocked users
		user.Status = Offline
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			break
		}
		blockedUsers = append(blockedUsers, user)

	}
	if err = rows.Err(); err != nil {
		log.Println("Error in returned rows: ", err)
	}

	return blockedUsers, err
}

func BlockUser(block Block) error {
	userID := block.UserID
	blockedID := block.BlockedID
	RemoveUserFromContacts(userID, blockedID)
	_, err := DB.Exec("INSERT INTO blocked (UserID, BlockedID) VALUES ($1, $2)", userID, blockedID)
	if err != nil {
		log.Panic(err)
	}
	return err
}

func UnblockUser(block Block) error {
	userID := block.UserID
	blockedID := block.BlockedID
	_, err := DB.Exec("DELETE FROM blocked WHERE UserID=$1 AND BlockedID=$2", userID, blockedID)
	if err != nil {
		log.Panic(err)
	}
	return err
}
func RandomBlock(nbMaxUser int) Block {
	return Block{
		UserID:    UserID(rand.Intn(nbMaxUser - 1)),
		BlockedID: UserID(rand.Intn(nbMaxUser - 1)),
	}
}
