package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/homocode/libro_iva_afip/api"
	db "github.com/homocode/libro_iva_afip/db"
	_ "github.com/lib/pq"
)

type config struct {
	DBDriver string
	DBSource string
}

func newDB(cfg config) *sql.DB {
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Can´t open connection to DB: ", err)
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal("Ping to DB fail, can´t connect to DB: ", pingErr)
	}

	return conn

}

func main() {
	conn := newDB(config{
		DBDriver: "postgres",
		DBSource: "postgres://localhost:5432/contable?sslmode=disable",
	})
	defer func() {
		_ = conn.Close()
		fmt.Println("DB close")
	}()

	store := db.NewStore(conn)
	s := api.NewServer(store, "localhost:8080")

	s.HttpServer.ListenAndServe()
}
