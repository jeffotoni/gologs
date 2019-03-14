// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
	"log"

	pg "github.com/jeffotoni/gologs/pkg/psql"
)

func InsertLog(jsonMsg string) bool {

	if len(jsonMsg) <= 0 {
		return false
	}

	Db := pg.Connect2()

	// var Db = pg.PostDb.Pgdb
	// // Db...
	// if interf := pg.Connect2(); interf != nil {
	// 	Db = interf.(*sql.DB)
	// } else {
	// 	return false
	// }

	///////////////////////////////////////////////////
	// Table gologs                                  //
	// CREATE TABLE gologs (                         //
	//     id serial not null primary key,           //
	//     time Timestamptz not null default  now(), //
	//     record Jsonb not null                     //
	// );                                            //
	///////////////////////////////////////////////////

	//data := time.Now().Format(cf.LayoutDate)
	//hora := time.Now().Format(cf.LayoutHour)

	insert := `INSERT INTO gologs(record)values($1)`
	_, err := Db.Exec(insert, jsonMsg)

	if err != nil {
		log.Println(err.Error())
		//log.Println(jsonMsg)
		return false
	}

	return true
}
