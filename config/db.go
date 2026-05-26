package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB config error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Database connected!")
}
