// Back-End in Go server
// @jeffotoni
// 2019-01-04

package postgres

import (
	"log"

	pg "github.com/jeffotoni/gologs/pkg/psql"
)

func Insert3Log(jsonMsg string) bool {

	if len(jsonMsg) <= 0 {
		return false
	}
	Db := pg.Connect2()
	insert := `INSERT INTO gologs(record)values($1)`
	_, err := Db.Exec(insert, jsonMsg)

	if err != nil {
		log.Println(err.Error())
		//log.Println(jsonMsg)
		return false
	}

	return true
}
