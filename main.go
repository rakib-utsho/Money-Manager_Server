package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"money-manager-server/config"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	// Tell the router: when someone visits "/", run this function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w = where you write your response back to the client
		// r = the incoming request (method, headers, body, etc.)
		fmt.Fprintln(w, "Money Manager API is running!")
	})

	// Start the server on port 8080
	// This line BLOCKS — it runs forever, listening for requests
	log.Println("Server Starting on port 8080....")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
