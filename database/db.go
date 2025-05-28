package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func InitDB() (*sql.DB, error) {
	// Connect to DB
	connStr, connStrErr := getDBConnectionString()
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
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName), err
}
