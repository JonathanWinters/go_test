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

const sqlFolderPath = "./Projects/go_test/sql"

/* --------------------------------- */
func CreateLevelTable() error {

	dirname, dirErr := os.UserHomeDir()
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	log.Printf("%s Home Directory:", dirname)

	path := filepath.Join(dirname, sqlFolderPath, "init.sql")

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

	//!TODO EFC: example of := for arg-use
	var dirname string
	dirname, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s Home Directory:", dirname)

	// path := filepath.Join(dirname, sqlFolderPath, "insert.sql")

	// //!INFO EFC: another scenario where just err can be used
	// c, ioErr := os.ReadFile(path)
	// if ioErr != nil {
	// 	log.Printf("ioErr \n")
	// 	err = ioErr
	// 	pk = -1
	// 	return
	// }
	// sqlQuery := string(c)

	sqlQuery := "INSERT INTO level (levelid, map, position, playerhitpoints) VALUES ($1, $2, $3, $4) RETURNING id;"

	jsonMap, _ := json.Marshal(level.Map)
	jsonPos, _ := json.Marshal(level.Position)

	err = DockerDb.db.QueryRow(sqlQuery, level.ID, jsonMap, jsonPos, level.PlayerHitPoints).Scan(&pk)

	return
}

func UpdateUserLevel(pk int, move int) (err string) {

	err = ""
	return
}
