package handlers

import (
	"encoding/json"
	"net/http"
)

type MultipartPayload struct {
	KeyPath     string `json:"keyPath"`
	ContentType string `json:"contentType"`
}

func CreateMultipartUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var mp MultipartPayload
	err := json.NewDecoder(r.Body).Decode(&mp)
	if err != nil {
		http.Error(w, "Invalid payload received "+err.Error(), 400)

		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mp)
}
