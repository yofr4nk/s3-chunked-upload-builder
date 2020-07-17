package main

import (
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/http/rest"
	"log"
	"net/http"
	"os"
)

func main() {
	r := rest.RouterHandler()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
