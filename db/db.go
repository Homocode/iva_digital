package db

import (
	"context"
	"database/sql"
	"log"
)

type Queries struct {
	db *sql.DB
}

func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func (q *Queries) PrepareStmt(ctx context.Context, query string) *sql.Stmt {
	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal("Can`t prepare statement. Error => ", err)
	}

	return stmt
}
