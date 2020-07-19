package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/http/handlers"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/uploading"
	"net/http"
)

// RouterHandler set the main config for routers
func RouterHandler(ufs *uploading.UploadFileService) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "pong")
	}).Methods("GET")

	router.HandleFunc("/get-multipart-upload", handlers.CreateMultipartUpload(ufs.CreateMultipartUpload)).Methods("POST")
	router.HandleFunc("/abort-multipart-upload", handlers.AbortMultipartUpload(ufs.AbortMultipartUpload)).Methods("POST")
	router.HandleFunc("/complete-multipart-upload", handlers.CompleteMultipartUpload(ufs.CompleteMultipartUpload)).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	return handler
}
