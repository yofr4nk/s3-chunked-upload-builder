package handlers

import (
	"encoding/json"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"net/http"
)

type MultipartBody struct {
	KeyPath     string `json:"keyPath"`
	ContentType string `json:"contentType"`
}

type createMultipartUpload func(filePathName string, contentType string) (domain.MultipartPayload, error)

func CreateMultipartUpload(createMultipartUpload createMultipartUpload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var mpb MultipartBody
		err := json.NewDecoder(r.Body).Decode(&mpb)
		if err != nil {
			http.Error(w, "Invalid payload received "+err.Error(), 400)

			return
		}

		mp, err := createMultipartUpload(mpb.KeyPath, mpb.ContentType)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mp)
	}
}
