package database

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

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

/* --------------------------------- */
func CreateLevelTable() error {

	path := filepath.Join("..", "..", "sql", "init.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		log.Printf("ioErr \n")
		return ioErr
	}
	sqlQuery := string(c)

	// log.Printf("%s", DockerDb)
	_, err := DockerDb.db.Exec(sqlQuery)
	if err != nil {
		log.Printf("err at DB.Exec \n")
		return err
	}

	return nil
}

func InsertLevel(level Level) (pk int, err error) {
	path := filepath.Join("..", "..", "sql", "insert.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		log.Printf("ioErr \n")
		err = ioErr
		pk = -1
		return
	}
	sqlQuery := string(c)

	jsonMap, _ := json.Marshal(level.Map)
	jsonPos, _ := json.Marshal(level.Position)

	err = DockerDb.db.QueryRow(sqlQuery, level.ID, jsonMap, jsonPos, level.PlayerHitPoints).Scan(&pk)

	return
}

func UpdateUserLevel(pk int, move int) (err string) {

	err = ""
	return
}
