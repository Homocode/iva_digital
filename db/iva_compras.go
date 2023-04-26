package db

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
)

func (q *Queries) CopyFromCsv(ctx context.Context, csvPath string) (sql.Result, error) {
	const copyQuery = `
	COPY comprabantes_compras_csv
	FROM '%s'
	DELIMITER ';'
	CSV HEADER;
	`
	query := fmt.Sprintf(copyQuery, csvPath)

	stmt := q.PrepareStmt(ctx, query)

	rs, err := stmt.ExecContext(ctx)

	return rs, err
}

func (q *Queries) CreateIvaComprasFromColumns(ctx context.Context, cols []string, cuitCliente string) (sql.Result, error) {
	var b bytes.Buffer
	for idx := range cols {
		// skip the first to columns name
		if idx == 0 || idx == 1 {
			continue
		}
		if idx == len(cols)-1 {
			b.WriteString(fmt.Sprintf("d.%q", cols[idx]))
		} else if idx != len(cols)-1 {
			b.WriteString(fmt.Sprintf("d.%q, ", cols[idx]))
		}
	}

	query := fmt.Sprintf("INSERT INTO iva_compras SELECT c.cuit, %s FROM clientes as c, comprabantes_compras_csv as d WHERE c.cuit = '%s';", &b, cuitCliente)
	fmt.Println("query ->>", query)
	rs, err := q.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
