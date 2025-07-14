package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateTables(connStr string) {
	db, err := sql.Open("postgres", connStr)

	CheckNil(err)
	CheckPing(err, db)

	CreateLevelTable(db)
}

// schema
/*
	User
		ID
		Levels: Array of Levels, each with an ID
	Level
		ID
		Map: number[x][y]
*/

func CreateLevelTable(db *sql.DB) {
	/*
		- ID
		- Map
		- PlayerHitPoints
	*/
	query := `CREATE TABLE IF NOT EXISTS level (
		id SERIAL PRIMARY KEY,
		map TEXT NOT NULL
		playerhitpoints
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	CheckNil(err)
}

func CreatePlayerTable() {
	/*
		- ID
		- Maps
	*/
}

func PingDB(connStr string) {

	db, err := sql.Open("postgres", connStr)

	CheckNil(err)
	CheckPing(err, db)

	defer db.Close()
}

func CheckNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPing(err error, db *sql.DB) {

	var checkErr = err

	if checkErr = db.Ping(); checkErr != nil {
		log.Fatal(err)
	}
}
