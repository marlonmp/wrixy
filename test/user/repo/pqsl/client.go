package pqsl_test

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var client *sql.DB

func init() {
	uri := "postgres://root:root@localhost:5432/wrixy?sslmode=disable"

	db, err := sql.Open("postgres", uri)

	if err != nil {
		panic(err)
	}

	client = db
}

func Client() *sql.DB {
	return client
}
