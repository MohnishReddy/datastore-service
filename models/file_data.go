package models

import "time"

type FileData struct {
	RefId             string
	FileId            string
	FileName          string
	CreatedAt         time.Time
	PrevServerDetails *ServerFileMeta
	NextServerDetails []*ServerFileMeta
	FileSize          int
	OriginalFileSize  int
	Data              []byte
}

type ServerFileMeta struct {
	ServerId  string
	ServerUrl string
	RefId     string
}
