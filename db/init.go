package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error
	// Open a connection to the SQLite database
	DB, err = sql.Open("sqlite3", "./daw-api.db")
	if err != nil {
		log.Fatal(err)
	}
}
