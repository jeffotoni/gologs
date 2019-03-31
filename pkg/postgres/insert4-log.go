// Back-End in Go server
// @jeffotoni
// 2019-01-04

package postgres

import (
  "database/sql"
  "fmt"
  "log"
  "strings"
  "time"

  "github.com/lib/pq"
)

type MessageDetailRecord struct {
  Record string
}

func Insert4Log(jsonMsg string) bool {

  // removendo aspas..
  DB_NAME = strings.Replace(DB_NAME, `"`, "", -1)
  DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)
  dbb, err := sql.Open(DB_SORCE, DBINFO)
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
