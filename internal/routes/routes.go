package routes

import (
	"net/http"

	"money-manager-server/internal/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/v1/health", handlers.HealthChecker)
}
