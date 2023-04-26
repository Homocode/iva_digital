package db

import (
	"context"
	"database/sql"
	"fmt"
)

func (q *Queries) getColumnsName(ctx context.Context, table string) ([]string, error) {
	columnsQuery := `
	SELECT ("column_name")
	FROM information_schema.columns
   	WHERE table_schema = 'public'
	AND table_name   = '%s';
	`
	query := fmt.Sprintf(columnsQuery, table)

	stmt := q.PrepareStmt(ctx, query)

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	var i string
	var items []string
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&i,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, err
}

func (q *Queries) deleteAllContent(ctx context.Context, table string) (sql.Result, error) {
	const deleteQuery = `
	TRUNCATE %s;
	`
	query := fmt.Sprintf(deleteQuery, table)

	stmt := q.PrepareStmt(ctx, query)

	rs, err := stmt.ExecContext(ctx)

	return rs, err
}
