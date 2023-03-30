package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	DBPath          string = "database/config.json"
	LocalhostDBPath string = "database/config_localhost.json"
	TestDBPath      string = "../../database/config_test.json"
)

var DB *sql.DB = nil

// const Connect = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", DB_username, DB_password, DB_ip, DB_name)

// Setup connects to database (DB should not be nil)  */
func Setup() {
	ConnectDB(LocalhostDBPath)
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

	_, err = DB.Exec("DO $$ BEGIN\n    CREATE TYPE status AS ENUM ('online', 'offline', 'busy', 'away');\nEXCEPTION\n    WHEN duplicate_object THEN null;\nEND $$;")

	if err != nil {
		log.Fatal("Error when building database CREATE TYPE status ", err)
	}
	fmt.Println("Successfully created or detected type status")

	initTable("users (id serial primary key, username text, status status )")
	initTable("messages (id serial primary key, userid integer, groupid integer, content text, time timestamp)")
	initTable("groups (id serial primary key, name text)")
	initTable("contacts (UserID integer, ContactID integer)")
	initTable("blocked (UserID integer, BlockedID integer)")
	initTable("usersToGroups (groupid integer, userid integer)")

}

func initTable(table string) {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " + table)
	if err != nil {
		log.Fatal("Error when building database CREATE TABLE IF NOT EXISTS "+table, err)
	}
	fmt.Println("Successfully created or detected table  " + table)
}
func DropAllTables() {
	tables := []string{"users", "messages", "groups", "contacts", "blocked", "usersToGroups"}
	for _, table := range tables {
		_, err := DB.Exec("DROP TABLE " + table)
		if err != nil {
			log.Fatal("Error when dropping table "+table, err)
		}
		fmt.Println("Successfully dropped table  " + table)
	}
}
