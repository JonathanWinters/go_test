package database

import (
	"database/sql"
	"log"

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

func ConnectDB(connStr string) (err error) {
	pgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error in sql.Open DB")
		return err
	}
	DockerDb.db = pgDB

	err = DockerDb.db.Ping()

	if err != nil {
		log.Println("Error in PINGING DB")
		log.Fatal(err)
		return err
	}
	log.Println("Connected to DB")
	return
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
