package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "api.db")

	if err != nil {
		panic("could not connect to db")
	}

	DB = db

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table" + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		desc TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table" + err.Error())
	}

	createEventRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id)
	)`

	_, err = DB.Exec(createEventRegistrationTable)

	if err != nil {
		panic("could not create registrations table" + err.Error())
	}
}
