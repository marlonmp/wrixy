package psql

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var client *sql.DB

func init() {

	uri := os.Getenv("PSQL_URI")

	db, err := sql.Open("postgres", uri)

	if err != nil {
		panic(err)
	}

	client = db
}

func Client() *sql.DB {
	return client
}
