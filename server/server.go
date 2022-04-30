package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10)
	mpf := r.MultipartForm

	for fileKey := range mpf.File {
		file, fileHeader, err := r.FormFile(fileKey)
		if err != nil {
			log.Fatalf("Failed to get file '%s' from MultipartForm\n", fileKey)
		}
		defer file.Close()
		fmt.Printf("Uploading '%s' to server...\n", fileHeader.Filename)

		path := "./server/files/" + fileHeader.Filename
		out, err := os.Create(path)
		if err != nil {
			log.Fatalf("Failed to open the path '%s'\n", path)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Successfully uploaded file(s) to server!")
}

func Entry() {
	// Need to make the ./server/files/ directory
	// with write permissions if it doesn't exist
	os.MkdirAll("./server/files", 0700)

	http.HandleFunc("/upload", handleUpload)
	http.ListenAndServe(":8080", nil)
}
