package main

import (
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/http/rest"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/loading"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/signing"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/storage/s3"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/uploading"
	"log"
	"net/http"
	"os"
)

func main() {
	envKeys, err := loading.GetEnvironmentKeys()
	if err != nil {
		log.Fatal(err)

		return
	}

	mediaStorage, err := s3.NewUploadStorage(envKeys.AwsAccessKey, envKeys.AwsSecretKey, envKeys.Bucket, envKeys.Region)
	if err != nil {
		log.Fatal(err)

		return
	}

	uploadFileService := uploading.NewUploadFileService(mediaStorage)
	uploadSignService := signing.NewSignUploadService(envKeys)

	r := rest.RouterHandler(uploadFileService, uploadSignService)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
