package router

import (
	"net/http"

	"money-manager-server/handlers"
)

const baseURL = "/api/v1"

func Setup(authHandler *handlers.AuthHandler) {
	http.HandleFunc(baseURL+"/user/register", authHandler.Register)
	http.HandleFunc(baseURL+"/user/login", authHandler.Login)
}
