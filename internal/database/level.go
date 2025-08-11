package database

import (
	"encoding/json"
	"log"

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

	sqlQuery := `CREATE TABLE IF NOT EXISTS "level"(id SERIAL PRIMARY KEY,levelid TEXT NOT NULL,map bytea NOT NULL,position bytea NOT NULL,playerhitpoints INT,created timestamp DEFAULT NOW());
`

	_, err := DockerDb.db.Exec(sqlQuery)
	if err != nil {
		log.Printf("err at DB.Exec \n")
		return err
	}

	return nil
}

func InsertLevel(level Level) (pk int, err error) {

	pk = -1

	sqlQuery := `INSERT INTO level (levelid, map, position, playerhitpoints) VALUES ($1, $2, $3, $4) RETURNING id`

	log.Printf("%s", sqlQuery)

	jsonMap, _ := json.Marshal(level.Map)
	jsonPos, _ := json.Marshal(level.Position)

	err = DockerDb.db.QueryRow(sqlQuery, level.ID, jsonMap, jsonPos, level.PlayerHitPoints).Scan(&pk)

	return
}
