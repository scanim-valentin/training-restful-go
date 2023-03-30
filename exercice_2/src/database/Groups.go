package database

import (
	"log"
	"service/utils"
)

// GroupID Unique users identifier
type GroupID int

// Group A group has a unique GroupID and a name (not unique).
type Group struct {
	ID   GroupID
	Name string
}

// CreateResponse is returned createGroup HTTP requests
type CreateResponse struct {
	ID   GroupID
	Name string
}

// NewGroup inserts new group and returns a unique newly created GroupID
func NewGroup(group Group) (GroupID, error) {
	name := group.Name
	var id GroupID
	err := DB.QueryRow("INSERT INTO groups (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

// NewGroup inserts new group and returns a unique newly created GroupID
func DeleteGroup(groupID GroupID) error {
	_, err := DB.Exec("DELETE FROM groups WHERE id=$1", groupID)
	if err != nil {
		log.Println("Couldn't delete group: ", err)
	}
	return err
}

func RandomGroup(id GroupID, nbMaxChar int) Group {
	return Group{
		ID:   id,
		Name: string(utils.RandomString(nbMaxChar)),
	}
}
