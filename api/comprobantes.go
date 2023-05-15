package api

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/homocode/libro_iva_afip/db"
	file "github.com/homocode/libro_iva_afip/files"
)

func (s *Server) UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath, err := file.UploadFile(r)
		if err != nil {
			return
		}
		// return result
		fmt.Fprintf(w, "Successfully Uploaded File")

		arg := db.LoadTxParams{
			CsvFilePath: filePath,
			CuitCliente: "20-30888725-4",
		}
		s.Store.LoadIvaComprasToDB(context.Background(), arg)
	}

}

func (s *Server) GetComprobantesForTable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
