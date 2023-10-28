package internal

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=secret password=secret host=localhost dbname=app-db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
