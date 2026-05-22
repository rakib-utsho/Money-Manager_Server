package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Money Manager API Running")
	})

	fmt.Println("Server running on :8000")

	http.ListenAndServe(":8000", nil)
}
