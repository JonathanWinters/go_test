package main

import (
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"
	server.SetHandlers()
	server.StartServer()
	database.ConnectDB(connStr)
	// database.PingDB(connStr)
	// database.CreateTables(connStr)
}
