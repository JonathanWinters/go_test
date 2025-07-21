package database

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
	_ "github.com/lib/pq"
)

type Level struct {
	ID               definitions.LevelID
	Map              data.Map
	OriginalPosition data.Positon
	PlayerHitPoints  int
}

/* --------------------------------- */
func CreateLevelTable() error {

	path := filepath.Join("..", "..", "sql", "init.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		// handle error.
		// log.Fatal(ioErr)
		log.Printf("ioErr \n")
		return ioErr
	}
	sqlQuery := string(c)
	log.Printf("%s", sqlQuery)

	// DB, openErr := sql.Open("postgres", data.DBConnectionString)
	// // util.CheckNil(err)
	// if openErr != nil {
	// 	return openErr
	// }
	// CheckPing(openErr, DB)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		log.Printf("err at DB.Exec \n")
		return err
	}

	return nil
}

func InsertLevel(level Level) int {
	query := `INSERT INTO level (levelid, map, originalposition, playerhitpoints)
		VALUES ($1, $2, $3, $4) RETURNING id`

	var pk int

	jsonMap, _ := json.Marshal(level.Map)
	jsonPos, _ := json.Marshal(level.OriginalPosition)

	// DB, openErr := sql.Open("postgres", data.DBConnectionString)
	// // util.CheckNil(err)
	// if openErr != nil {
	// 	return -1
	// }
	// CheckPing(openErr, DB)

	err := DB.QueryRow(query, level.ID, jsonMap, jsonPos, level.PlayerHitPoints).Scan(&pk)
	util.CheckNil(err)

	return pk
}

func GetLevelByID() {

}

func UpdateUserLevel() {

}
