package domain

type MultipartPayload struct {
	Bucket   string
	UploadId string
	KeyPath  string
}

type MultipartCompletedInput struct {
	UploadId       string         `json:"uploadId"`
	KeyPath        string         `json:"keyPath"`
	CompletedParts CompletedParts `json:"completedParts"`
}

type CompletedPart struct {
	ETag       string `json:"etag"`
	PartNumber int64  `json:"partNumber"`
}

type CompletedParts []CompletedPart
