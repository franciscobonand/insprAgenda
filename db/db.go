package db

import (
	"database/sql"

	// lib for postgre usage
	_ "github.com/lib/pq"
)

// ConnectWithDB returns a reference for the SQL database
func ConnectWithDB() *sql.DB {
	connection := "user=postgres dbname=insprTaskManager " +
		"password=postgres123 " +
		"host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
