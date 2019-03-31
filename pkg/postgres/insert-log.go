// Back-End in Go server
// @jeffotoni
// 2019-01-04

package postgres

import (
	"database/sql"
	"log"
)

func InsertLog(jsonMsg string) bool {

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
	_, err := Db.Exec(insert, jsonMsg)

	if err != nil {
		log.Println(err.Error())
		//log.Println(jsonMsg)
		return false
	}

	return true
}
