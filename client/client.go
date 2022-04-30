package client

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const addr = "localhost:8080"

func Entry(filePath string) {
	err := upload(filePath)
	if err != nil {
		log.Fatalln(err)
	}
}

func upload(filePath string) error {
	sanitizedPath := strings.ReplaceAll(filePath, " ", "")
	paths := strings.Split(sanitizedPath, ",")

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer func() {
			// m.Close() is important so the requset knows the boundary
			m.Close()
			w.Close()
		}()
		for i, path := range paths {
			f, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			fileKey := fmt.Sprintf("file%d", i)
			if fw, err := m.CreateFormFile(fileKey, f.Name()); err != nil {
				return
			} else {
				if _, err = io.Copy(fw, f); err != nil {
					return
				}
			}
		}
	}()

	url := fmt.Sprintf("http://%s/upload", addr)
	req, _ := http.NewRequest("POST", url, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	fmt.Printf("Successfully uploaded %d file(s)!\n", len(paths))
	return nil
}
