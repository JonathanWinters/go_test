package main

import (
	"log"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/server"
	_ "github.com/lib/pq"
)

func main() {

	err := database.ConnectDB(data.DBConnectionString)
	if err != nil {
		log.Printf("Err in Connecting to DB")
		return
	}
	server.SetHandlers()
	server.StartServers()
}
