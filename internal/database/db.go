package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// InitializeDB creates a connection to the PostgreSQL database and returns the *sql.DB object.
func InitDB() *sql.DB {
	// Retrieve database connection details from environment variables
	host := os.Getenv("DATABASE_HOST") // Use "postgres" as the service name in Docker
	port := os.Getenv("DB_PORT")       // Default to "5432" if not specified
	user := os.Getenv("POSTGRES_USER") // Default "root"
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Construct the connection string using environment variables
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, password, host, port)

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
