package database

import (
	"database/sql"
	"log"
)

type Member struct {
	UserID  UserID
	GroupID GroupID
}

func AddUserToGroup(member Member) error {
	groupID := member.GroupID
	userID := member.UserID
	_, err := DB.Exec("INSERT INTO usersToGroups (groupid, UserID) VALUES ($1, $2)", groupID, userID)
	return err
}

func RemoveUserFromGroup(member Member) error {
	groupID := member.GroupID
	userID := member.UserID
	_, err := DB.Exec("DELETE FROM usersToGroups WHERE userid=$1 AND groupid=$2", userID, groupID)
	return err
}

func GetUserGroups(userid UserID) ([]Group, error) {
	var groups []Group
	rows, err := DB.Query("SELECT id, name FROM groups INNER JOIN usersToGroups ON groupid=id AND userid=$1 ORDER BY id ASC", userid)
	if err != nil {
		log.Println("Error when getting groups: ", err)
		return groups, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic("Failed to close rows in GetGroup: ", err)
		}
	}(rows)

	// Reading rows
	for rows.Next() {
		var group Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			break
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error in rows: ", err)
	}
	return groups, err
}

func GetUsersInGroup(groupid GroupID) ([]User, error) {
	var users []User
	// TODO handle blocked users (no name, disconnected status)
	rows, err := DB.Query("SELECT id, username, status FROM users INNER JOIN usersToGroups ON userid=id AND groupid=$1 ORDER BY id ASC", groupid)
	if err != nil {
		log.Println("Error when getting users: ", err)
		return users, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic("Failed to close rows in GetGroup: ", err)
		}
	}(rows)

	// Reading rows
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Status); err != nil {
			break
		}
		// fmt.Println("USER ", user)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error in rows: ", err)
	}
	return users, err
}
