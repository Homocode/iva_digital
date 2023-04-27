package api

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	db "github.com/homocode/libro_iva_afip/db"
	file "github.com/homocode/libro_iva_afip/files"
)

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("fileUp.html")
		t.Execute(w, nil)
	} else {
		filePath, err := file.UploadFile(r)
		// TODO: handle this err
		if err != nil {
			return
		}
		// return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")

		arg := db.LoadTxParams{
			CsvFilePath: filePath,
			CuitCliente: "20-38",
		}
		s.store.LoadIvaComprasToDB(context.Background(), arg)
	}
}
