package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() error {
	// Connect to DB
	connStr, connStrErr := getDBConnectionString()
	if connStrErr != nil {
		return connStrErr
	}
	log.Println("Connection String Established")

	db, sqlConnErr := sql.Open("postgres", connStr)
	if sqlConnErr != nil {
		return sqlConnErr
	}

	// Test the connection
	if pingErr := db.Ping(); pingErr != nil {
		return pingErr
	}

	Db = db
	log.Println("Connected to Postgres")
	return nil
}

func getDBConnectionString() (string, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Access environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName), err
}
