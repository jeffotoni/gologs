// Back-End in Go server
// @jeffotoni
// 2019-01-04

package maps

import (
	"log"
	"sync"

	"github.com/jeffotoni/gologs/pkg/postgres"
	"github.com/jeffotoni/gologs/pkg/redis"
)

var m sync.Map

func Save(key int, value string) bool {

	m.Store(key, value)
	//v, ok := m.Load(k)
	//m.Delete(k)
	return true
}

func SavePg() {
	//erase map
	m.Range(func(key interface{}, value interface{}) bool {
		// save in DB
		postgres.Insert5Log(value.(string))
		m.Delete(key)
		// time.Sleep(time.Millisecond * 1500)
		return true
	})

	log.Println("fim save Postgres!")
}

func SaveRedis() {

	//erase map
	m.Range(func(key interface{}, value interface{}) bool {
		// save in DB
		redis.SaveRedis(key.(int), value.(string))
		m.Delete(key)
		// time.Sleep(time.Millisecond * 1500)
		return true
	})

	log.Println("fim save Redis!")
}
