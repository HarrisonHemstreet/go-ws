package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitializeDB creates a connection to the PostgreSQL database and returns the *sql.DB object.
func InitDB() *sql.DB {
	// Database connection string
	connStr := "user=root dbname=root password=root host=127.0.0.1 port=5440 sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Verify the connection with a ping
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	return db
}
