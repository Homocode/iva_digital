package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/homocode/libro_iva_afip/api"
	db "github.com/homocode/libro_iva_afip/db"
	_ "github.com/lib/pq"
)

func main() {
	type config struct {
		DBDriver string
		DBSource string
	}

	cfg := config{
		DBDriver: "postgres",
		DBSource: "postgres://localhost:5432/contable?sslmode=disable",
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Can´t open connection to DB: ", err)
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal("Ping to DB fail, can´t connect to DB: ", pingErr)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, "localhost:8080")

	http.HandleFunc("/upload", server.UploadFile)

	server.Start()
}
