package handlers

import (
	"encoding/json"
	"net/http"
)

type abortMultipartBody struct {
	KeyPath  string `json:"keyPath"`
	UploadId string `json:"uploadId"`
}

type abortMultipartUpload func(filePathName string, uploadId string) error

func AbortMultipartUpload(abortMultipartUpload abortMultipartUpload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var amp abortMultipartBody
		err := json.NewDecoder(r.Body).Decode(&amp)
		if err != nil {
			http.Error(w, "Invalid payload received "+err.Error(), 400)

			return
		}

		err = abortMultipartUpload(amp.KeyPath, amp.UploadId)
		if err != nil {
			http.Error(w, "something went wrong aborting upload "+err.Error(), 500)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
