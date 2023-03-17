package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB = nil

// const Connect = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", DB_username, DB_password, DB_ip, DB_name)

// Setup connects to database (DB should not be nil)  */
func Setup() {
	ConnectDB("database/config.json")
}

// Close should be called before closing service */
func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Fatal: Failed to close database:", DB)
	}
	fmt.Println("Successfully disconnected from database")
}

// ConnectDB connects and setup database environment based on json config file found at path
// You should only use it for testing purpose, otherwise use Setup for regular API use
func ConnectDB(path string) {
	// Let's first read the `config.json` file
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	payload := map[string]string{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	connect := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", payload["username"], payload["password"], payload["ip"], payload["name"])
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("Error while opening DB: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error while trying to reach DB: ", err)
	}
	fmt.Println("Ping OK -> ", connect)
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, username text, ip inet, port int )")
	if err != nil {
		log.Fatal("Error when building database (CREATE TABLE IF NOT EXISTS users (id serial primary key, username text, ip inet, port int )) ",err)
	}
	fmt.Println("Successfully created or detected table 'user'")
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, source integer, destination integer, content text, time timestamp)")
	if err != nil {
		log.Fatal("Error when building database CREATE TABLE IF NOT EXISTS messages (id serial primary key, source integer, destination integer, content text, time timestamp) ",err)
	}
	fmt.Println("Successfully created or detected table 'messages'")
}
