package entities

import "time"

type MongoFileData struct {
	RefId             string            `bson:"ref_id"`
	FileId            string            `bson:"file_id"`
	FileName          string            `bson:"file_name"`
	CreatedAt         time.Time         `bson:"created_at"`
	PrevServerDetails *ServerFileMeta   `bson:"prev_server_details,omitempty"`
	NextServerDetails []*ServerFileMeta `bson:"next_server_details,omitempty"`
	FileSize          int               `bson:"file_size"`
	OriginalFileSize  int               `bson:"original_file_size"`
	Data              []byte            `bson:"data"`
}

type ServerFileMeta struct {
	ServerId  string `bson:"server_id,omitempty"`
	ServerUrl string `bson:"server_url,omitempty"`
	RefId     string `bson:"ref_id,omitempty"`
}
