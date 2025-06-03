package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

// EFFECT: Initializes the test database connection on localhost.
// REQUIRES: Assumes that the .env file is present in the root directory
func InitTestDB() (*sql.DB, error) {
	// Connect to DB
	connStr, connStrErr := getTestDBConnectionString()
	if connStrErr != nil {
		return nil, connStrErr
	}
	log.Println("Connection String Established")

	db, sqlConnErr := sql.Open("postgres", connStr)
	if sqlConnErr != nil {
		return nil, sqlConnErr
	}

	// Test the connection
	if pingErr := db.Ping(); pingErr != nil {
		return nil, pingErr
	}

	log.Println("Connected to Postgres")
	return db, nil
}

func getTestDBConnectionString() (string, error) {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Access environment variables
	testUser := os.Getenv("TEST_DB_USER")
	testPassword := os.Getenv("TEST_DB_PASSWORD")
	testName := os.Getenv("TEST_DB_NAME")

	return fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		testUser, testPassword, testName), err
}
