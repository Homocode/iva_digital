package api

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

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

		s.store.CopyFromCsv(context.Background(), fmt.Sprintf("./csv/%s", tempFile.Name()))

	}
}
