package models

type FileDownloadResponse struct {
	FileBytes         *[]byte          `json:"file_bytes"`
	FileId            string           `json:"file_id"`
	FileName          string           `json:"file_name"`
	NextServerDetails *[]ServerDetails `json:"next_server_details,omitempty"`
}

type ServerDetails struct {
	ServerId  string `json:"server_details"`
	ServerUrl string `json:"server_url"`
	RefId     string `json:"ref_id"`
}
