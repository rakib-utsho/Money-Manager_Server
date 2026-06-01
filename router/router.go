package router

import (
	"fmt"
	"net/http"

	"money-manager-server/handlers"
	"money-manager-server/middleware"
)

const baseURL = "/api/v1"

func Setup(authHandler *handlers.AuthHandler) {
	http.HandleFunc(baseURL+"/user/register", authHandler.Register)
	http.HandleFunc(baseURL+"/user/login", authHandler.Login)

	http.HandleFunc(baseURL+"/test-auth", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.UserIDKey).(int)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "You are authenticated", "user_id": ` + fmt.Sprintf("%d", userId) + ` }`))
	}))
}
