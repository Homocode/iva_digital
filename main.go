package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/homocode/libro_iva_afip/api"
	db "github.com/homocode/libro_iva_afip/db"
	util "github.com/homocode/libro_iva_afip/utils"
	_ "github.com/lib/pq"
)

func newDB(cfg util.Config) *sql.DB {
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Can´t open connection to DB: ", err)
	}

	time.AfterFunc(2*time.Second, func() {
		pingErr := conn.Ping()
		if pingErr != nil {
			log.Fatal("Ping to DB fail, can´t connect to DB: ", pingErr)
		}
	})

	return conn

}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can`t load configuration enviroment variables")
	}

	fmt.Println(config)

	conn := newDB(config)
	defer func() {
		_ = conn.Close()
		fmt.Println("DB close")
	}()

	store := db.NewStore(conn)
	s := api.NewServer(store, config.ServerAddress)

	s.HttpServer.ListenAndServe()
}
