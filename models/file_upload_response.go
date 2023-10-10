package models

type FileUploadResponse struct {
	FileId          string `json:"file_id"`
	FileName        string `json:"file_name"`
	FileSize        int    `json:"file_size"`
	RefId           string `json:"ref_id"`
	DataUploaded    int    `json:"data_uploaded"`
	CompressionUsed bool   `json:"compression_used"`
}
