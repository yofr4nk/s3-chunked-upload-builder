package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"log"
)

type UploadStorage struct {
	session *session.Session
	bucket  string
}

func NewUploadStorage(accessKey string, secretKey string, bucket string, region string) (*UploadStorage, error) {
	sess, err := CreateAwsSession(accessKey, secretKey, region)
	if err != nil {
		return nil, err
	}

	return &UploadStorage{
		session: sess,
		bucket:  bucket,
	}, nil
}

func (storage *UploadStorage) CreateMultipartUpload(filePathName string, contentType string) (domain.MultipartPayload, error) {
	svc := s3.New(storage.session)

	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(storage.bucket),
		Key:         aws.String(filePathName),
		ContentType: aws.String(contentType),
	}

	mpOutput, err := svc.CreateMultipartUpload(input)
	if err != nil {
		log.Print(err.Error())

		return domain.MultipartPayload{}, err
	}

	return domain.MultipartPayload{
		Bucket:   *mpOutput.Bucket,
		UploadId: *mpOutput.UploadId,
		KeyPath:  *mpOutput.Key,
	}, nil
}

func (storage *UploadStorage) AbortMultipartUpload(filePathName string, uploadId string) error {
	svc := s3.New(storage.session)

	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   aws.String(storage.bucket),
		Key:      aws.String(filePathName),
		UploadId: aws.String(uploadId),
	}

	_, err := svc.AbortMultipartUpload(abortInput)
	if err != nil {
		log.Print(err.Error())

		return err
	}

	return nil
}

func (storage *UploadStorage) CompleteMultipartUpload(cmi domain.MultipartCompletedInput) (string, error) {
	svc := s3.New(storage.session)

	completedMultipartUploadInput := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(storage.bucket),
		Key:      aws.String(cmi.KeyPath),
		UploadId: aws.String(cmi.UploadId),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: fillCompletedParts(cmi.CompletedParts),
		},
	}

	completeResponse, err := svc.CompleteMultipartUpload(completedMultipartUploadInput)
	if err != nil {
		return "", err
	}

	return completeResponse.String(), nil
}

func fillCompletedParts(completedParts domain.CompletedParts) []*s3.CompletedPart {
	var cParts []*s3.CompletedPart

	for i := 0; i < len(completedParts); i++ {
		cp := s3.CompletedPart{
			ETag:       &completedParts[i].ETag,
			PartNumber: &completedParts[i].PartNumber,
		}

		cParts = append(cParts, &cp)
	}

	return cParts
}
