package main

import (
	"github.com/JonathanWinters/go_test/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	database.PingDB(connStr)
	database.CreateTables(connStr)
}
