package db

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func New(url string) *sql.DB {
	db, err := sql.Open("postgres", url)

	// set max connection
	maxConnection, _ := strconv.Atoi(os.Getenv("RDS_MAX_CONN"))
	db.SetMaxOpenConns(maxConnection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
