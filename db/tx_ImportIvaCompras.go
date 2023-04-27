package db

import (
	"context"
	"database/sql"
)

type LoadTxParams struct {
	CsvFilePath string
	CuitCliente string
}

type LoadTxResult struct {
	copyFileResult         sql.Result
	createIvaComprasResult sql.Result
}

func (s *SQLStore) LoadIvaComprasToDB(ctx context.Context, arg LoadTxParams) (LoadTxResult, error) {
	var rs LoadTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		// Copy the content from the csv files to the comprobantes_compras_csv table
		rs.copyFileResult, err = s.CopyFromCsv(ctx, arg.CsvFilePath)
		if err != nil {
			return err
		}

		// Populate the table iva_compras with the select columns
		// from comprobantes_compras_csv and associate cuit of the cliente
		rs.createIvaComprasResult, err = s.InsertComprobantes(ctx, arg.CuitCliente)
		if err != nil {
			return err
		}

		_, err = q.deleteAllContent(ctx, "comprabantes_compras_csv")
		if err != nil {
			return err
		}

		return nil
	})

	return rs, err
}
