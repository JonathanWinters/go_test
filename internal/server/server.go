package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Start the server and listen on port 8080
	fmt.Println("Server starting on port 5442...")
	err := http.ListenAndServe(":5442", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
