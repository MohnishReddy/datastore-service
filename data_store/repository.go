package data_store

import (
	"context"
	"datastore-service/data_store/mongo"
	"datastore-service/models"
)

type Datastore interface {
	StoreFile(ctx context.Context, fileData *models.FileData) error
	ReadFile(ctx context.Context, refId string) (fileData *models.FileData, err error)
}

var repository Datastore

func InitDataStore(databaseName string, collectionName string) error {
	repo, err := mongo.NewMongoDataStore(databaseName, collectionName)
	if err != nil {
		return err
	}

	repository = repo
	return nil
}

func GetRepository() Datastore {
	return repository
}
