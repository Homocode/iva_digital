package api

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	db "github.com/homocode/libro_iva_afip/db"
)

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("fileUp.html")
		t.Execute(w, nil)
	} else {
		// parse input
		r.ParseMultipartForm(10 << 20)

		// retrieve file
		file, _, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		// write temporary file on our server
		tempFile, err := os.CreateTemp("./csv", "upload-*.csv")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return
		}
		tempFile.Write(buf.Bytes())

		// return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")

		absPath, err := filepath.Abs(tempFile.Name())
		// TODO: handle this err
		if err != nil {
			return
		}

		arg := db.LoadTxParams{
			CsvFilePath: absPath,
			CuitCliente: "20-38",
		}
		fmt.Println(">>", arg.CsvFilePath)
		s.store.LoadIvaComprasToDB(context.Background(), arg)
	}
}
