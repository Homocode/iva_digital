package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Copia el contenido del archivo csv (comprobantes de compras) importado de AFIP a la
// tabla comprobantes_compras_csv
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

// Inserta los comprobantes de pago de la tabla comprobantes_compras_csv y
// el cuit del cliente asociado a los comprobantes.
func (q *Queries) InsertComprobantes(ctx context.Context, cuitCliente string) (sql.Result, error) {
	const insertQuery = `
	INSERT INTO comprobantes_compras
	SELECT c.cuit, d.* 
	FROM clientes as c, comprabantes_compras_csv as d 
	WHERE c.cuit = '%s';
	`

	query := fmt.Sprintf(insertQuery, cuitCliente)

	rs, err := q.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (q *Queries) ListComprobantes(ctx context.Context) {

}
