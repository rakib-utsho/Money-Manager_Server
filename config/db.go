package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DB is the shared connection pool used across the entire app
var DB *sql.DB

func ConnectDB() {
	// build connection string from environment variables
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error

	// sql.Open doesn't actually connect yet
	// it just prepares the connection pool
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB config error:", err)
	}

	// Ping actually attempts the connection
	// if this fails, app exits immediately at startup
	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Database connected!")
}
