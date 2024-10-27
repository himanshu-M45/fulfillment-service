package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Database is an interface that abstracts the database operations.
// 1. Ping() error: Pings the database.
// 2. Query(query string, args ...interface{}) (*sql.Rows, error): Executes a query that returns rows.
// 3. QueryRow(query string, args ...interface{}) *sql.Row: Executes a query that returns a single row.
// 4. Exec(query string, args ...interface{}) (sql.Result, error): Executes a query that doesn't return rows.
type Database interface {
	Ping() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var db *sql.DB

// InitDB initializes the database connection.
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Successfully connected to database!")
}

func GetDB() *sql.DB {
	return db
}
