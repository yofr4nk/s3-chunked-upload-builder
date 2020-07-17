package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/http/handlers"
	"net/http"
)

// RouterHandler set the main config for routers
func RouterHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "pong")
	}).Methods("GET")

	router.HandleFunc("/get-multipart-upload", handlers.CreateMultipartUpload).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	return handler
}
