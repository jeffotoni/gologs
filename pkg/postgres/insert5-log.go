// Back-End in Go server
// @jeffotoni
// 2019-01-04

package postgres

import (
	"database/sql"
	"log"
)

func Insert5Log(jsonMsg string) bool {

	if len(jsonMsg) <= 0 {
		return false
	}

	var Db = PostDb.Pgdb
	// Db...
	if interf := Connect(); interf != nil {
		Db = interf.(*sql.DB)
	} else {
		return false
	}

	insert := `INSERT INTO gologs(record)values($1)`
	tx, err := Db.Begin()
	if err != nil {
		log.Println(err)
		return false
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare(insert)
	if err != nil {
		log.Println(err)
		return false
	}

	defer stmt.Close() // danger!

	_, err = stmt.Exec(jsonMsg)
	if err != nil {
		log.Println(err)
		return false
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
