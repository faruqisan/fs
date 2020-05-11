package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {

	var (
		path string
	)

	flag.StringVar(&path, "p", "", "-p=path/to/folder")
	flag.Parse()

	if path == "" {
		log.Fatal("no path provided")
	}

	// Simple static webserver:
	log.Println("file server running at", port, "serving path:", path)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
