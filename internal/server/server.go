package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Start the server and listen on port 8080
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
