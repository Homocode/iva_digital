package db

import (
	"context"
	"database/sql"
	"fmt"
)

func (q *Queries) deleteAllContent(ctx context.Context, table string) (sql.Result, error) {
	const deleteQuery = `
	TRUNCATE %s;
	`
	query := fmt.Sprintf(deleteQuery, table)

	stmt := q.PrepareStmt(ctx, query)

	rs, err := stmt.ExecContext(ctx)

	return rs, err
}
