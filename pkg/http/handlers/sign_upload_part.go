package handlers

import (
	"encoding/json"
	"github.com/yofr4nk/aws-signed-request-sigv4/sign"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"net/http"
)

type signUploadPart func(signingUpload domain.UploadSign) sign.SignatureResponse

func SignUploadPart(signUploadPart signUploadPart) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var uploadSign domain.UploadSign
		err := json.NewDecoder(r.Body).Decode(&uploadSign)
		if err != nil {
			http.Error(w, "Invalid payload received "+err.Error(), 400)

			return
		}

		uploadPartSigned := signUploadPart(uploadSign)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(uploadPartSigned)
	}
}
