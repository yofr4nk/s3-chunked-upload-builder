package signing

import (
	"fmt"
	"github.com/yofr4nk/aws-signed-request-sigv4/sign"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"time"
)

type SignUploadService struct {
	config domain.EnvironmentConfig
}

func NewSignUploadService(config domain.EnvironmentConfig) *SignUploadService {
	return &SignUploadService{config: config}
}

func (su SignUploadService) SignUploadPart(uploadSign domain.UploadSign) sign.SignatureResponse {
	pts := sign.PayloadToSing{
		Date:        time.Now().UTC(),
		KeyPath:     uploadSign.KeyPath,
		QueryString: fmt.Sprintf("partNumber=%s&uploadId=%s", uploadSign.PartNumber, uploadSign.UploadId),
		ContentMD5:  uploadSign.ContentMD5,
		Method:      "PUT",
	}
	config := sign.ConfigData{
		Bucket:          su.config.Bucket,
		SecretAccessKey: su.config.AwsSecretKey,
		AccessKeyId:     su.config.AwsAccessKey,
		Region:          su.config.Region,
	}

	return sign.CalculateSignature(pts, config)
}
