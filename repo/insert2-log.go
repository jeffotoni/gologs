// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	pg "github.com/jeffotoni/gologs/pkg/psql"
)

func Insert2Log(jsonMsg string) bool {

	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// log.Println(m.Alloc)
	// s := make([]byte, 1024*1024)
	// s[0] = 0

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

	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
	Db.SetConnMaxLifetime(0)
	defer Db.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(jsonMsg string) {
		defer wg.Done()
		//_, err := db.Exec(msg)
		insert := `INSERT INTO gologs(record)values($1)`
		_, err = Db.Exec(insert, jsonMsg)
		if err != nil {
			log.Println(err)
		}
	}(jsonMsg)
	wg.Wait()

	return true
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

	// if m.Alloc > 104857600 { // 100Mb
	// 	log.Println(m.Alloc)
	// 	time.Sleep(time.Second * 1)
	// }

	// insert := `INSERT INTO gologs(record)values($1)`
	// //insert := `INSERT INTO gologs(record)values('` + jsonMsg + `')`
	// _, err = Db.Exec(insert, jsonMsg)
	// //_, err = Db.Exec(insert)

	// // runtime.ReadMemStats(&m)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	//log.Println(jsonMsg)
	// 	return false
	// }

	// return true
}
