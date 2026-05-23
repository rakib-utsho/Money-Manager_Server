package handlers

import (
	"fmt"
	"net/http"
)

func HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Api is healthy")
}
