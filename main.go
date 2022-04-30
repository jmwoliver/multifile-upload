package main

import (
	"flag"
	"log"

	"github.com/jmwoliver/multifile-upload/client"
	"github.com/jmwoliver/multifile-upload/server"
)

var (
	filePath string
	entry    string
)

func main() {

	flag.StringVar(&entry, "entry", "", "Specify to start 'server' or 'client'.")
	flag.StringVar(&filePath, "files", "", "The files to upload. ex. --files='file1.png, file2.png, file3.png'")
	flag.Parse()

	switch entry {
	case "server":
		if filePath != "" {
			log.Fatalf("Cannot use --files flag when spawning server.")
		}
		server.Entry()
	case "client":
		if filePath == "" {
			log.Fatalln("The --files flag must be used when spawning client. ex. --files='file1.png, file2.png, file3.png'")
		}
		client.Entry(filePath)
	default:
		log.Fatalf("Not a valid value for --entry. Must be 'server' or 'client'.")
	}
}
