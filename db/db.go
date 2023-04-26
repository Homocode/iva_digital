package db

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db DBTX
}

func NewQueries(db DBTX) *Queries {
	return &Queries{db: db}
}

func (q *Queries) PrepareStmt(ctx context.Context, query string) *sql.Stmt {
	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal("Can`t prepare statement. Error => ", err)
	}

	return stmt
}
