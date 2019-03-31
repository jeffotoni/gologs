// Back-End in Go server
// @jeffotoni
// 2019-01-04

package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
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
	DB_NAME = strings.Replace(DB_NAME, `"`, "", -1)
	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)
	Db, err := sql.Open(DB_SORCE, DBINFO)
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

}
