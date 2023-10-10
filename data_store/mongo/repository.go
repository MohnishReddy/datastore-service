package mongo

import (
	"context"
	"datastore-service/data_store/mongo/entities"
	"datastore-service/pkg"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	db *mongo.Database
	c  *mongo.Collection
}

type MongoDb interface {
	InsertFileDataAndGetRefId(ctx context.Context, fileData *entities.MongoFileData) error
	FindFileDataFromRefId(ctx context.Context, refId string) (fileData *entities.MongoFileData, err error)
}

func NewMongoRepo(databaseName string, collectionName string) (MongoDb, error) {
	cfg := pkg.GetConfig()

	if cfg.Database == nil {
		fmt.Println("Not connected to database!")
		return nil, nil
	}

	// Set MongoDB connection options.
	mongoServerUrl := cfg.Database.DBUrl
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoServerUrl)

	// Connect to MongoDB.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	// Get a handle to the database and collection.
	database := client.Database(databaseName)
	collection := database.Collection(collectionName)

	return &repository{
		db: database,
		c:  collection,
	}, nil
}

func (r *repository) InsertFileDataAndGetRefId(ctx context.Context, fileData *entities.MongoFileData) error {
	_, err := r.c.InsertOne(ctx, fileData, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindFileDataFromRefId(ctx context.Context, refId string) (fileData *entities.MongoFileData, err error) {
	data := r.c.FindOne(ctx, bson.M{"ref_id": refId})
	if data.Err() != nil {
		return nil, errors.New("file data not found for provided refId")
	}

	fileData = &entities.MongoFileData{}
	err = data.Decode(fileData)
	if err != nil {
		return nil, err
	}

	return
}
