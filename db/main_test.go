package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	testDb, err := sql.Open("postgres", "postgres://localhost:5432/contable?sslmode=disable")
	if err != nil {
		log.Fatal("Can`t connect to DB", err)
	}

	testQueries = NewQueries(testDb)

	os.Exit(m.Run())
}
