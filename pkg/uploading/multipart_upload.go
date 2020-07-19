package uploading

import (
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
)

type UploadFileRepository interface {
	CreateMultipartUpload(filePathName string, contentType string) (domain.MultipartPayload, error)
	AbortMultipartUpload(filePathName string, uploadId string) error
	CompleteMultipartUpload(completedMultipartInput domain.MultipartCompletedInput) (string, error)
}

type UploadFileService struct {
	repository UploadFileRepository
}

func NewUploadFileService(repository UploadFileRepository) *UploadFileService {
	return &UploadFileService{repository: repository}
}

func (ufs UploadFileService) CreateMultipartUpload(filePathName string, contentType string) (domain.MultipartPayload, error) {
	mp, err := ufs.repository.CreateMultipartUpload(filePathName, contentType)
	if err != nil {
		return domain.MultipartPayload{}, err
	}

	return mp, nil
}

func (ufs UploadFileService) AbortMultipartUpload(filePathName string, uploadId string) error {
	err := ufs.repository.AbortMultipartUpload(filePathName, uploadId)
	if err != nil {
		return err
	}

	return nil
}

func (ufs UploadFileService) CompleteMultipartUpload(completedMultipartInput domain.MultipartCompletedInput) (string, error) {
	cr, err := ufs.repository.CompleteMultipartUpload(completedMultipartInput)
	if err != nil {
		return "", err
	}

	return cr, nil
}
