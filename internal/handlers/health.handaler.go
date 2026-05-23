package handlers

import (
	"net/http"

	"money-manager-server/internal/utils"
)

func HealthChecker(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(
		w,
		http.StatusOK,
		true,
		"Api is Healthy",
		nil,
	)
}
