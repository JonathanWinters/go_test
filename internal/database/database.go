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
	//!FIXED
	// !INFO EFC: typically we always want to inspect the error before assigning values returned (this scenario is ok since the db is nillable)
	DockerDb.db = pgDB

	//!FIXED
	// !INFO EFC: use defers to gracefully close objects if any error occurred (prevent memory leaks)
	// !INFO EFC: we can't do it here because it will close when this func returns
	// defer DockerDb.db.Close()
	err = DockerDb.db.Ping()

	// log.Printf("%s", DockerDb)
	if err != nil {
		log.Println("Error in PINGING DB")
		log.Fatal(err)
		return err
	}
	log.Println("Connected to DB")
	return
}

// !INFO EFC: in our world, the pkey is usually the roundid or userid, which is included in client requests
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
