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

type UploadSign struct {
	KeyPath    string `json:"keyPath"`
	PartNumber string `json:"partNumber"`
	UploadId   string `json:"uploadId"`
	ContentMD5 string `json:"contentMD5"`
}
