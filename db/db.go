package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // the underscore defines a blank import
	// it shows that this package is only used for its side effects adn should not be removed
	// during build process
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db") // here the first argument is the driver, and the second is the database
 
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	DB.SetMaxOpenConns(10) //this says that at most 10 connections can be open at a time
	DB.SetMaxIdleConns(5)  // this says at least 5 connections can be idle at a time

	createTable()
}

func createTable() {

	// this is a raw sql query
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable) // this executes the query

	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	// this is a raw sql query
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable) // this executes the query

	if err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}
}
