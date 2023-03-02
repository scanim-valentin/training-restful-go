package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB = nil

// const Connect = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", DB_username, DB_password, DB_ip, DB_name)

/*
* Setup (connect to database)
 */
func Setup() {

	// Let's first read the `config.json` file
	content, err := os.ReadFile("database/config.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	fmt.Println(json.Valid(content))
	// Now let's unmarshall the data into `payload`
	payload := map[string]string{}
	fmt.Print(payload)
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	fmt.Print(payload)
	connect := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", payload["username"], payload["password"], payload["ip"], payload["name"])
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to ", connect)
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, username text, ip inet, port int )")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created or detected table 'user'")
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, source integer, destination integer, content text, time timestamp )")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created or detected table 'messages'")
}

/*
* Should be called before closing service
 */
func Close() {
	log.Fatal(DB.Close())
	fmt.Println("Successfully disconnected from database")
}
