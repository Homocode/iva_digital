package db

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
)

func (q *Queries) CopyFromCsv(ctx context.Context, csvPath string) (sql.Result, error) {
	const copyQuery = `
	COPY csv_data
	FROM '%s'
	DELIMITER ';'
	CSV HEADER;
	`
	query := fmt.Sprintf(copyQuery, csvPath)

	stmt := q.PrepareStmt(ctx, query)

	rs, err := stmt.ExecContext(ctx)

	return rs, err
}

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

func (q *Queries) CreateIvaCompras(ctx context.Context, cuitCliente string) (sql.Result, error) {
	cols, err := q.getColumnsName(ctx, "iva_compras")
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	for idx := range cols {
		if idx == 0 {
			continue
		}
		if idx == len(cols)-1 {
			b.WriteString(fmt.Sprintf("d.%q", cols[idx]))
		} else if idx != len(cols)-1 {
			b.WriteString(fmt.Sprintf("d.%q, ", cols[idx]))
		}
	}

	query := fmt.Sprintf("INSERT INTO iva_compras SELECT c.cuit, %s FROM clientes as c, csv_data as d WHERE c.cuit = '%s';", &b, cuitCliente)
	fmt.Println("query ->>", query)
	rs, err := q.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
