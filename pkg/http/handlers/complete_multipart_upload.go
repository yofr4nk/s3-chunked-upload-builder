package handlers

import (
	"encoding/json"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"net/http"
)

type CompleteUploadResponse struct {
	Data string `json:"data"`
}

type completeMultipartUpload func(multipartCompletedInput domain.MultipartCompletedInput) (string, error)

func CompleteMultipartUpload(completeMultipartUpload completeMultipartUpload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var mci domain.MultipartCompletedInput
		err := json.NewDecoder(r.Body).Decode(&mci)
		if err != nil {
			http.Error(w, "Invalid payload received "+err.Error(), 400)

			return
		}

		completedResponse, err := completeMultipartUpload(mci)
		if err != nil {
			http.Error(w, err.Error(), 500)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CompleteUploadResponse{
			Data: completedResponse,
		})
	}
}
