// Back-End in Go server
// @jeffotoni
// 2019-01-04

package psql

///
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once

///////////////////////////////////////////////////
// Table gologs                                  //
// CREATE TABLE gologs (                         //
//     id serial not null primary key,           //
//     time Timestamptz not null default  now(), //
//     record Jsonb not null                     //
// );                                            //
///////////////////////////////////////////////////

///
/////// DATA BASE
var (
	DB_NAME     = os.Getenv("DB_NAME")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_SSL      = os.Getenv("DB_SSL")
	DB_SORCE    = os.Getenv("DB_SORCE")

	///api
	API_ENV     = ""
	API_HOST    = ""
	HOST_CONFIG = ""
)

type PgStruct struct {
	Pgdb *sql.DB
}

type StatusMsg struct {
	Msg     string `json:msg`
	Db      string `json:db`
	Uuidzip string `json:uuidzip`
}

// cache sync.Map
type cache struct {
	mm sync.Map
	sync.Mutex
}

var (
	err    error
	PostDb PgStruct
)

var (
	pool = &cache{}
)

func init() {
	if len(os.Getenv("DB_PORT")) <= 0 {
		DB_PORT = "5432"
	}
	if len(os.Getenv("DB_SSL")) <= 0 {
		DB_SSL = "disable"
	}
	if len(os.Getenv("DB_SORCE")) <= 0 {
		DB_SORCE = "postgres"
	}

	if len(os.Getenv("DB_NAME")) <= 0 {
		fmt.Printf("\033[0;31m")
		println(" Error, export DB_NAME is permited!")
		fmt.Printf("\033[0;0m")
		showEnvDb()
		return
	}

	if len(os.Getenv("DB_HOST")) <= 0 {
		fmt.Printf("\033[0;31m")
		println(" Error, export DB_HOST is permited!")
		fmt.Printf("\033[0;0m")
		showEnvDb()
		return
	}

	if len(os.Getenv("DB_USER")) <= 0 {
		fmt.Printf("\033[0;31m")
		println(" Error, export DB_USER is permited!")
		fmt.Printf("\033[0;0m")
		showEnvDb()
		return
	}

	if len(os.Getenv("DB_PASSWORD")) <= 0 {
		fmt.Printf("\033[0;31m")
		println(" Error, export DB_PASSWORD is permited!")
		println("\033[0m")
		showEnvDb()
		return
	}
}

func showEnvDb() {
	println(" Please set the environment variable for the postgres database!")
	fmt.Printf("\033[0;33m")
	println("   Info:")
	println("    - DB_HOST=your-host")
	println("    - DB_NAME=your-name")
	println("    - DB_USER=your-user")
	println("    - DB_PASSWORD=xxxxxx")
	println("    - DB_PORT=5432")
	println("\033[0m")
}

// put sync.Map
func (c *cache) put(key, value interface{}) {

	c.Lock()
	defer c.Unlock()
	c.mm.Store(key, value)
}

// get sync.Map
func (c *cache) get(key interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	v, _ := c.mm.Load(key)
	return v

}

// setLoad... fn func() interface{}
func (c *cache) loadStore(key interface{}, fc func() (interface{}, error)) (interface{}, error) {

	c.Lock()
	defer c.Unlock()

	if v, ok := c.mm.Load(key); ok {
		return v, nil
	}

	// treat or error
	val, err := fc()

	// error return
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	c.mm.Store(key, val)
	return val, nil
}

// conectando de forma segura usando goroutine
func Connect() interface{} {

	if dbPg := pool.get(DB_NAME); dbPg != nil {

		// return objeto conexao
		return dbPg.(*sql.DB)

	} else {

		// removendo aspas..
		DB_NAME = strings.Replace(DB_NAME, `"`, "", -1)

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

		// func for execute
		// in loadStore
		// when two or more
		// goroutine at moment
		fn := func() (interface{}, error) {

			once.Do(func() {
				PostDb.Pgdb, err = sql.Open(DB_SORCE, DBINFO)
			})

			if err != nil {
				log.Println(err.Error())
				return nil, err
			}

			if ok2 := PostDb.Pgdb.Ping(); ok2 != nil {
				log.Println("connect error...: ", ok2)
				return nil, err
			}

			//log.Println("connect return sucess:: client [" + DB_NAME + "]")
			return PostDb.Pgdb, nil
		}

		// get connect
		// load cache loadStore
		sqlDb, err := pool.loadStore(DB_NAME, fn)

		if err != nil {
			// error
			return nil
		}

		if sqlDb != nil {
			return sqlDb.(*sql.DB)

		} else {
			return nil
		}
	}
}

// conectando de forma segura usando goroutine
func Connect2() interface{} {

	if PostDb.Pgdb != nil {
		// return objeto conexao
		return PostDb.Pgdb

	} else {

		// removendo aspas..
		DB_NAME = strings.Replace(DB_NAME, `"`, "", -1)

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

		once.Do(func() {
			PostDb.Pgdb, err = sql.Open(DB_SORCE, DBINFO)
		})

		if err != nil {
			log.Println(err.Error())
			return err
		}

		if ok2 := PostDb.Pgdb.Ping(); ok2 != nil {
			log.Println("connect error...: ", ok2)
			return err
		}

		//log.Println("connect return sucess:: client [" + DB_NAME + "]")
		return PostDb.Pgdb
	}
}
