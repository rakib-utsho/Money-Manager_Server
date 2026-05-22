package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	Database "money-manager-server/internal/database"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	err = Database.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Money Manager API Running")
	})

	fmt.Println("Server running on :8000")

	http.ListenAndServe(":8000", nil)
}
