// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strings"

	pg "github.com/jeffotoni/gologs/pkg/psql"
)

func Insert2Log(jsonMsg string) bool {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Println(m.Alloc)
	s := make([]byte, 1024*1024)
	s[0] = 0

	if len(jsonMsg) <= 0 {
		return false
	}

	// removendo aspas..
	pg.DB_NAME = strings.Replace(pg.DB_NAME, `"`, "", -1)
	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.DB_HOST, pg.DB_PORT, pg.DB_USER, pg.DB_PASSWORD, pg.DB_NAME, pg.DB_SSL)
	Db, err := sql.Open(pg.DB_SORCE, DBINFO)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	//Db.SetMaxOpenConns(2)
	//Db.SetMaxIdleConns(1)
	//Db.SetConnMaxLifetime(time.Second * 10)
	defer Db.Close()

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

	// insert := `INSERT INTO gologs(record)values($1)`
	insert := `INSERT INTO gologs(record)values('` + jsonMsg + `')`
	//_, err = Db.Exec(insert, jsonMsg)
	_, err = Db.Exec(insert)

	runtime.ReadMemStats(&m)
	log.Println(m.Alloc)

	return true

	if err != nil {
		log.Println(err.Error())
		//log.Println(jsonMsg)
		return false
	}

	return true
}
