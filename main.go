package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // registers postgres driver, we don't call it directly

	"money-manager-server/config"
	"money-manager-server/handlers"
	"money-manager-server/repository"
	"money-manager-server/router"
	"money-manager-server/services"
)

func main() {

	// load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	config.ConnectDB()

	userRepo := &repository.UserRepository{DB: config.DB}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthService: authService}

	router.Setup(authHandler)

	// Tell the router: when someone visits "/", run this function
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// w = where you write your response back to the client
	// 	// r = the incoming request (method, headers, body, etc.)
	// 	fmt.Fprintln(w, "Money Manager API is running!")
	// })

	// Start the server on port 8080
	// This line BLOCKS — it runs forever, listening for requests
	log.Println("Server Starting on port 8080....")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
