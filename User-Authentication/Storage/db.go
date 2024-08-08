package storage

import (
    "log"

    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() {
    var err error
    connStr := "user=youruser dbname=yourdb sslmode=disable" 
    db, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        log.Fatalln(err)
    }
}
