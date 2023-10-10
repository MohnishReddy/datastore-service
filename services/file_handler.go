package services

import (
	"context"
	"datastore-service/models"
	"time"

	"github.com/google/uuid"
)

type fileUploadHandler interface {
	HandleFileUpload(ctx context.Context, fileId string, fileName string, file []byte, compressionReq bool) (refId *string, dataUploaded int, err error)
}

type fileDownloadHandler interface {
	HandleFileDownload(ctx context.Context, refId string) (file *models.FileData, err error)
}

type fileStruct struct {
	prevServerDetails *models.ServerFileMeta
}

func NewFileUploadHandler(prevServerUrl string, prevServerId string, prevDataRefId string) fileUploadHandler {
	return &fileStruct{
		prevServerDetails: &models.ServerFileMeta{
			ServerId:  prevServerId,
			ServerUrl: prevServerUrl,
			RefId:     prevDataRefId,
		},
	}
}

func NewFileDownloadHandler() fileDownloadHandler {
	return &fileStruct{}
}

func (f *fileStruct) HandleFileUpload(ctx context.Context, fileId string, fileName string, file []byte, compressionReq bool) (*string, int, error) {
	finalFile := &file
	fileSize := len(file)
	newFileSize := len(file)
	if compressionReq {
		compressedFile, err := NewDataHandler().Compress(&file)
		if err != nil {
			return nil, 0, err
		}

		finalFile = compressedFile
		newFileSize = len(*compressedFile)
	}

	refId := uuid.New().String()

	fileData := &models.FileData{
		RefId:            refId,
		FileId:           fileId,
		FileName:         fileName,
		OriginalFileSize: fileSize,
		FileSize:         newFileSize,
		CreatedAt:        time.Now(),
		PrevServerDetails: &models.ServerFileMeta{
			ServerId:  f.prevServerDetails.ServerId,
			ServerUrl: f.prevServerDetails.ServerUrl,
			RefId:     f.prevServerDetails.RefId,
		},
		Data: *finalFile,
	}

	err := repo.StoreFile(ctx, fileData)
	if err != nil {
		return nil, 0, err
	}

	return &refId, newFileSize, nil
}

func (f *fileStruct) HandleFileDownload(ctx context.Context, refId string) (*models.FileData, error) {
	fileData, err := repo.ReadFile(ctx, refId)
	if err != nil {
		return nil, err
	}

	if fileData.OriginalFileSize != fileData.FileSize {
		decompressedFile, err := NewDataHandler().Decompress(&fileData.Data)
		if err != nil {
			return nil, err
		}

		fileData.Data = *decompressedFile

	}

	return fileData, nil
}
