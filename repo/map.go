// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
	"log"
	"sync"
	"time"
)

var m sync.Map

func Map(key int, value string) bool {

	m.Store(key, value)
	//v, ok := m.Load(k)
	//m.Delete(k)
	return true
}

func SavePg() {

	//erase map
	m.Range(func(key interface{}, value interface{}) bool {
		// save in DB
		// Insert5Log(value.(string))
		m.Delete(key)
		time.Sleep(time.Millisecond * 1500)
		return true
	})

	log.Println("fim save Postgres!")
}
