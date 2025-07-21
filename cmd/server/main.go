package main

import (
	"log"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/server"
	_ "github.com/lib/pq"
)

func main() {

	database.ConnectDB(data.DBConnectionString)

	createErr := database.CreateLevelTable()
	if createErr != nil {
		log.Printf("Table was NOT created")
		log.Fatal(createErr)
		return
	}
	log.Printf("Table WAS created")

	server.SetHandlers()
	server.StartServer()
}
