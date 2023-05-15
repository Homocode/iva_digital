package file

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(r *http.Request) (string, error) {
	// parse input
	r.ParseMultipartForm(10 << 20)

	// retrieve file
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	// write temporary file on the server
	tempFile, err := os.CreateTemp("", "upload-comprobantes-compras.--*.csv")
	if err != nil {
		return "", err
	}
	absPath, err := filepath.Abs(tempFile.Name())
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}
	tempFile.Write(buf.Bytes())

	return absPath, nil
}
