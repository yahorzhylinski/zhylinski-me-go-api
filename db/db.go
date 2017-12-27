package db

import (
  "sync"
  "database/sql"
  "gopkg.in/reform.v1"
  "gopkg.in/reform.v1/dialects/postgresql"
  _ "github.com/lib/pq"
  "os"
  "log"
)

var instance *reform.DB
var once sync.Once

var DATABASE_ADAPTER string = "postgres"
var CONNECTION_STRING string = "postgres://hsyalyec:Ln4V-O9rvzXvovYD-IKjkOkmVCW_PdxY@packy.db.elephantsql.com:5432/hsyalyec"

func DB() *reform.DB {
  once.Do(func() {
    connection, err := sql.Open(DATABASE_ADAPTER, CONNECTION_STRING)
    if err != nil {
      panic(err)
    }

    logger := log.New(os.Stderr, "SQL: ", log.Flags())
    instance = reform.NewDB(connection, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
  })

  return instance
}
