package mongo

import (
	"context"
	"datastore-service/data_store/mongo/entities"
	"datastore-service/models"
	"errors"
)

type MongoDataStore struct {
	repo MongoDb
}

func NewMongoDataStore(databaseName string, collectionName string) (*MongoDataStore, error) {
	mongoRepo, err := NewMongoRepo(databaseName, collectionName)
	if err != nil {
		return nil, err
	}

	return &MongoDataStore{
		repo: mongoRepo,
	}, nil
}

func (mds *MongoDataStore) StoreFile(ctx context.Context, fileData *models.FileData) error {
	err := mds.repo.InsertFileDataAndGetRefId(ctx, mapFileDataModelToEntity(fileData))
	return err
}

func (mds *MongoDataStore) ReadFile(ctx context.Context, refId string) (fileData *models.FileData, err error) {
	fileDataEntity, err := mds.repo.FindFileDataFromRefId(ctx, refId)
	if err != nil {
		return nil, err
	}

	if fileDataEntity == nil {
		return nil, errors.New("file for provided ref id not found")
	}
	return mapFileDataEntityToModel(fileDataEntity), nil
}

func mapFileDataModelToEntity(fileData *models.FileData) *entities.MongoFileData {
	var prevServerDetails *entities.ServerFileMeta

	if fileData.PrevServerDetails != nil {
		prevServerDetails = &entities.ServerFileMeta{
			ServerId:  fileData.PrevServerDetails.ServerId,
			ServerUrl: fileData.PrevServerDetails.ServerUrl,
			RefId:     fileData.PrevServerDetails.RefId,
		}
	}

	var nextServerDetails []*entities.ServerFileMeta
	for _, serverDetail := range fileData.NextServerDetails {
		nextServerDetails = append(nextServerDetails, &entities.ServerFileMeta{
			ServerId:  serverDetail.ServerId,
			ServerUrl: serverDetail.ServerUrl,
			RefId:     serverDetail.RefId,
		})
	}

	return &entities.MongoFileData{
		RefId:             fileData.RefId,
		FileId:            fileData.FileId,
		FileName:          fileData.FileName,
		CreatedAt:         fileData.CreatedAt,
		PrevServerDetails: prevServerDetails,
		NextServerDetails: nextServerDetails,
		FileSize:          fileData.FileSize,
		OriginalFileSize:  fileData.OriginalFileSize,
		Data:              fileData.Data,
	}
}

func mapFileDataEntityToModel(fileData *entities.MongoFileData) *models.FileData {
	var prevServerDetails *models.ServerFileMeta

	if fileData.PrevServerDetails != nil {
		prevServerDetails = &models.ServerFileMeta{
			ServerId:  fileData.PrevServerDetails.ServerId,
			ServerUrl: fileData.PrevServerDetails.ServerUrl,
			RefId:     fileData.PrevServerDetails.RefId,
		}
	}

	var nextServerDetails []*models.ServerFileMeta
	for _, serverDetail := range fileData.NextServerDetails {
		nextServerDetails = append(nextServerDetails, &models.ServerFileMeta{
			ServerId:  serverDetail.ServerId,
			ServerUrl: serverDetail.ServerUrl,
			RefId:     serverDetail.RefId,
		})
	}

	return &models.FileData{
		RefId:             fileData.RefId,
		FileId:            fileData.FileId,
		FileName:          fileData.FileName,
		CreatedAt:         fileData.CreatedAt,
		PrevServerDetails: prevServerDetails,
		NextServerDetails: nextServerDetails,
		FileSize:          fileData.FileSize,
		OriginalFileSize:  fileData.OriginalFileSize,
		Data:              fileData.Data,
	}
}
