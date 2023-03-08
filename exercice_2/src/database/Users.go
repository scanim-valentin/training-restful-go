package database

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"service/utils"
)

// UserID Unique user identifier
type UserID int

// User A user has a unique UserID, a name (not unique), an IP address and a Port.
// IP and Port are also used to indicate the user's status (connected or not)
type User struct {
	ID   UserID
	Name string
	IP   net.IP
	Port int
}

// LoginResponse Login response is returned to login and register HTTP requests
type LoginResponse struct {
	ID       UserID
	Username string
	UserList []User
}

// InsertNewUser Insert new user and returns newly created UserID
func InsertNewUser(name string, ip net.IP, port string) UserID {
	var id UserID
	err := DB.QueryRow("INSERT INTO users (username, ip, port) VALUES ($1, $2, $3) RETURNING id", name, ip.String(), port).Scan(&id)
	if err != nil {
		log.Panic(err)
	}
	return id
}

func LoginUser(ip net.IP, port string, id UserID) *[]byte {
	name := make([]byte, 0)

	switch err := DB.QueryRow("UPDATE users SET ip = $1, port = $2 WHERE id = $3 RETURNING username", ip.String(), fmt.Sprint(port), id).Scan(&name); err {
	case nil:
		break // gross
	case sql.ErrNoRows:
		return nil
	default:
		log.Fatal("Unhandled error from database: ", err)
	}
	return &name
}

// GetUserList Retrieve online user list
func GetUserList() []User {
	// SQL Queries
	// Order must be specified to match test specification
	// There is no default order (see https://stackoverflow.com/questions/6585574/postgres-default-sort-by-id-worldship "Rows are returned in an unspecified order")
	rows, err := DB.Query("SELECT * FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}(rows)

	var users []User

	// Reading rows
	for rows.Next() {
		var user User
		var ipAux string
		if err := rows.Scan(&user.ID, &user.Name, &ipAux, &user.Port); err != nil {
			return users
		}
		user.IP = net.ParseIP(ipAux)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Panic(err)
	}
	return users
}

func GetAddress(destination UserID) (net.IP, int) {
	rows, err := DB.Query("SELECT ip,port FROM users WHERE id = $1", destination)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var otherIpStr string
	var otherPort int
	// Reading rows
	fmt.Print(rows)
	rows.Next()
	if err = rows.Scan(&otherIpStr, &otherPort); err != nil {
		log.Panic(err)
	}
	var otherIp net.IP
	if otherIp = net.ParseIP(otherIpStr); otherIp == nil {
		log.Panic("IP address could not be read for user ", destination)
	}
	return otherIp, otherPort
}

func SetStatusOffline(id UserID) {
	_, err := DB.Exec("UPDATE users SET ip = $1 WHERE id = $2", utils.IPUnspecified.String(), id)
	if err != nil {
		log.Fatal(err)
	}
}

func RandomUser(id UserID, nbMaxChar int) User {
	return User{
		ID:   id,
		Name: string(utils.RandomString(nbMaxChar)),
		IP:   net.IPv4(1, 1, 1, 1),
		Port: 1111,
	}
}
