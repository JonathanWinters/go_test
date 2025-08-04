package database

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	_ "github.com/lib/pq"
)

type Level struct {
	ID              definitions.LevelID
	Map             data.Map
	Position        data.Positon
	PlayerHitPoints int
}

const sqlFolderPath = "./Projects/go_test/sql"

/* --------------------------------- */
func CreateLevelTable() error {

	// load JSON from disk
	file, fileErr := os.Open("init.sql")
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	// Create a byte slice to store the read data
	buffer := make([]byte, 1024) // Read in chunks of 1024 bytes

	for {
		// Read from the file into the buffer
		_, err := file.Read(buffer)
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	}

	sqlQuery := string(buffer)

	_, err := DockerDb.db.Exec(sqlQuery)
	if err != nil {
		log.Printf("err at DB.Exec \n")
		return err
	}

	return nil
}

func InsertLevel(level Level) (pk int, err error) {

	pk = -1
	// load JSON from disk
	homeDir, _ := os.UserHomeDir()
	file, err := os.Open(homeDir + "/Projects/go_test/sql/insert.sql")

	log.Printf("%s", homeDir)
	if err != nil {
		pk = -1
		return
	}
	defer file.Close()

	// Create a byte slice to store the read data
	buffer := make([]byte, 1024) // Read in chunks of 1024 bytes

	for {
		// Read from the file into the buffer
		_, err := file.Read(buffer)
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	}

	sqlQuery := string(buffer)

	log.Printf("%s", sqlQuery)

	jsonMap, _ := json.Marshal(level.Map)
	jsonPos, _ := json.Marshal(level.Position)

	err = DockerDb.db.QueryRow(sqlQuery, level.ID, jsonMap, jsonPos, level.PlayerHitPoints).Scan(&pk)

	return
}
