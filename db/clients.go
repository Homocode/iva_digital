package db

import (
	"context"
)

type CreateClienteParams struct {
	Cuit        string `db:"cuit" json:"cuit"`
	RazonSocial string `db:"razon_social" json:"razon_social"`
}

func (q *Queries) CreateCliente(ctx context.Context, args CreateClienteParams) (Clientes, error) {
	const createQuery = `
	INSERT INTO clientes (
		cuit,
		razon_social
	) VALUES (
		$1, $2
	)
	RETURNING *
	`

	stmt := q.PrepareStmt(ctx, createQuery)

	row := stmt.QueryRowContext(ctx, args.Cuit, args.RazonSocial)
	var i Clientes
	err := row.Scan(
		&i.Cuit,
		&i.RazonSocial,
	)

	return i, err
}

func (q *Queries) ListClientes(ctx context.Context) ([]Clientes, error) {
	const getQuery = `
	SELECT * FROM clientes;
	`
	stmt := q.PrepareStmt(ctx, getQuery)

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Clientes{}
	for rows.Next() {
		var i Clientes
		if err := rows.Scan(
			&i.Cuit,
			&i.RazonSocial,
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
	return items, nil
}
