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
	host := getEnv("DATABASE_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "5440")
	user := getEnv("POSTGRES_USER", "root")
	password := getEnv("POSTGRES_PASSWORD", "root")
	dbname := getEnv("POSTGRES_DB", "postgres")

	// Construct the connection string using environment variables
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, password, host, port)
	fmt.Println("connStr: ", connStr)

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

// getEnv retrieves the value of the environment variable named by the key
// or returns the default value if the variable is not set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
