package database

import (
	"database/sql"
	"log"

	"github.com/JonathanWinters/go_test/internal/data/dummydata"
	"github.com/JonathanWinters/go_test/internal/util"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(connStr string) error {
	DB, err := sql.Open("postgres", connStr)
	// util.CheckNil(err)
	if err != nil {
		return err
	}
	CheckPing(err, DB)
	return nil
}

func CreateTables(connStr string) {

	CreateLevelTable()

	ogPosition := util.FindIndex2DArray(dummydata.Map, 4)

	dummyLevel := Level{dummydata.LevelID, dummydata.Map, ogPosition, dummydata.PlayerHitPoints}
	InsertLevel(dummyLevel)
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
		log.Printf("Error in PINGING DB")
		log.Fatal(err)
	}
}
