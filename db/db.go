package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() { //WHAT IS THE STRUCTURE I AM USING IN HERE? WHAT IS IT FOR? LOOK FOR IT!!!!!!!!!!!!
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		fmt.Println("ERROR AT INITDB")
		panic("Could not connect to a database.")
	}

	DB.SetMaxOpenConns(10) //Max amount of connections that can be opened
	DB.SetMaxIdleConns(5)  //Connections open if no one is even using a connection

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL, 
		dateTime TEXT NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create tables.")
	}
}
