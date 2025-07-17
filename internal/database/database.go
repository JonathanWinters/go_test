package database

import (
	"database/sql"
	"log"

	"github.com/JonathanWinters/go_test/internal/data/dummydata"
	"github.com/JonathanWinters/go_test/internal/util"
	_ "github.com/lib/pq"
)

func ConnectDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	util.CheckNil(err)
	CheckPing(err, db)
	return db
}

func CreateTables(connStr string) {
	db, err := sql.Open("postgres", connStr)

	util.CheckNil(err)
	CheckPing(err, db)

	CreateLevelTable(db)

	ogPosition := util.FindIndex2DArray(dummydata.Map, 4)

	dummyLevel := Level{dummydata.LevelID, dummydata.Map, ogPosition, dummydata.PlayerHitPoints}
	InsertLevel(db, dummyLevel)
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

func CreatePlayerTable() {
	/*
		- ID
		- Maps
	*/
}

func PingDB(connStr string) {

	db, err := sql.Open("postgres", connStr)

	util.CheckNil(err)
	CheckPing(err, db)

	defer db.Close()
}

func CheckPing(err error, db *sql.DB) {

	var checkErr = err

	if checkErr = db.Ping(); checkErr != nil {
		log.Fatal(err)
	}
}
