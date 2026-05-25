package routes

import (
	"net/http"

	"money-manager-server/internal/handlers"

)

const apiBase = "/api/v1"

func RegisterRoutes() {
	http.HandleFunc(apiBase+"/health", handlers.HealthChecker)
	http.HandleFunc(apiBase+"/users/register", handlers.RegisterUserHandler)
}

