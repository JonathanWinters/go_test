package database

import (
	"database/sql"
	"log"

	"github.com/JonathanWinters/go_test/internal/data/dummydata"
	"github.com/JonathanWinters/go_test/internal/util"
	_ "github.com/lib/pq"
)

type db struct {
	db *sql.DB
}

type LevelRow struct {
	Id  int
	Map []byte
}

var DockerDb db

func ConnectDB(connStr string) error {
	pgDB, err := sql.Open("postgres", connStr)

	DockerDb.db = pgDB

	if err != nil {
		log.Printf("Error in sql.Open DB")
		return err
	}

	checkErr := DockerDb.db.Ping()

	// log.Printf("%s", DockerDb)
	if checkErr != nil {
		log.Printf("Error in PINGING DB")
		log.Fatal(err)
		return checkErr
	}

	return nil
}

func UpdateLevelHPAndPositionByPrimaryKey(pk int, hp int, pos []byte) error {

	sqlQuery := `UPDATE "level" 
				SET position = $1, 
					playerhitpoints = $2 
				WHERE id = $3`

	_, err := DockerDb.db.Exec(sqlQuery, pos, hp, pk)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func CreateTables(connStr string) {

	CreateLevelTable()

	ogPosition := util.FindIndex2DArray(dummydata.Map, 4)

	dummyLevel := Level{dummydata.LevelID, dummydata.Map, ogPosition, dummydata.PlayerHitPoints}
	InsertLevel(dummyLevel)
}

func GetMapByPrimaryKey(pk int) (levelMap []byte, err error) {

	sqlQuery := `SELECT map FROM "level" WHERE id = $1`

	err = DockerDb.db.QueryRow(sqlQuery, pk).Scan(&levelMap)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func GetPlayerHitPointsByPrimaryKey(pk int) (hitpoints int, err error) {

	sqlQuery := `SELECT playerhitpoints FROM "level" WHERE id = $1`

	err = DockerDb.db.QueryRow(sqlQuery, pk).Scan(&hitpoints)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func GetPositionByPrimaryKey(pk int) (pos []byte, err error) {

	sqlQuery := `SELECT position FROM "level" WHERE id = $1`

	err = DockerDb.db.QueryRow(sqlQuery, pk).Scan(&pos)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func CheckPing(err error) {

	var checkErr = err

	if checkErr = DockerDb.db.Ping(); checkErr != nil {
		log.Printf("Error in PINGING DB")
		log.Fatal(err)
	}
}
