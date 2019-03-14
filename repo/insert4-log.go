// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
  "database/sql"
  "fmt"
  "log"
  "strings"
  "time"

  pg "github.com/jeffotoni/gologs/pkg/psql"
  "github.com/lib/pq"
)

type MessageDetailRecord struct {
  Record string
}

func Insert4Log(jsonMsg string) bool {

  // removendo aspas..
  pg.DB_NAME = strings.Replace(pg.DB_NAME, `"`, "", -1)
  DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    pg.DB_HOST, pg.DB_PORT, pg.DB_USER, pg.DB_PASSWORD, pg.DB_NAME, pg.DB_SSL)
  dbb, err := sql.Open(pg.DB_SORCE, DBINFO)
  if err != nil {
    log.Println(err.Error())
    return false
  }

  dbb.SetMaxIdleConns(10)
  dbb.SetMaxOpenConns(10)
  dbb.SetConnMaxLifetime(0)
  a := time.Now()
  txn, err := dbb.Begin()
  if err != nil {
    log.Fatal(err)
  }

  lenjson := len(jsonMsg)

  stmt, _ := txn.Prepare(pq.CopyIn("messagedetailrecord", "record")) // MessageDetailRecord is the table name
  m := &MessageDetailRecord{
    Record: jsonMsg,
  }
  mList := make([]*MessageDetailRecord, 0, lenjson)
  for i := 0; i < lenjson; i++ {
    // fmt.Println(i)
    mList = append(mList, m)
  }

  fmt.Println(m)

  for _, user := range mList {
    _, err := stmt.Exec(string(user.Record))
    if err != nil {
      log.Println(err)
      continue
    }
  }
  _, err = stmt.Exec()
  if err != nil {
    log.Fatal(err)
  }
  err = stmt.Close()
  if err != nil {
    log.Fatal(err)
  }
  err = txn.Commit()
  if err != nil {
    log.Fatal(err)
  }
  delta := time.Now().Sub(a)
  fmt.Println(delta.Nanoseconds())
  fmt.Println("Program finished successfully")

  return true

}
