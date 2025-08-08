package main

import (
	"log"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/database"
	"github.com/JonathanWinters/go_test/internal/server"
	_ "github.com/lib/pq" //!INFO EFC: this annoying quirk creates package-level variables and executes the init() function of that package
)

func main() {

	//!FIXED
	//!INFO EFC: err vars
	err := database.ConnectDB(data.DBConnectionString)
	if err != nil {
		log.Printf("Err in Connecting to DB")
		return
	}
	//!FIXED
	//!INFO EFC: db creation handled via docker/scripts
	server.SetHandlers()
	server.StartServers()
}
