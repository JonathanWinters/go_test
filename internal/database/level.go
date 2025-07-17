package database

import (
	"database/sql"

	"github.com/JonathanWinters/go_test/internal/data"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
)

type Level struct {
	ID               definitions.LevelID
	Map              data.Map
	OriginalPosition data.Positon
	PlayerHitPoints  int
}

/* --------------------------------- */
func CreateLevelTable(db *sql.DB) {
	/*
		- ID
		- Map
		- PlayerHitPoints
	*/
	query := `CREATE TABLE IF NOT EXISTS level (
		id SERIAL PRIMARY KEY,
		map TEXT NOT NULL
		originalposition
		playerhitpoints
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	util.CheckNil(err)
}

func InsertLevel(db *sql.DB, level Level) int {
	query := `INSERT INTO level (map, originalposition, playerhitpoints)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, level.Map, level.OriginalPosition, level.PlayerHitPoints).Scan(&pk)
	util.CheckNil(err)
	return pk
}

func GetLevelByID() {

}

func UpdateUserLevel() {

}
