package database

import "math/rand"

const (
	Online  string = "online"
	Offline string = "offline"
	Busy    string = "busy"
	Away    string = "away"
)

func RandomStatus() string {
	statusArray := [4]string{Online, Offline, Busy, Away}
	return statusArray[rand.Intn(len(statusArray))]
}
