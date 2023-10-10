package services

import (
	"datastore-service/data_store"
	"sync"
)

var repo data_store.Datastore
var muLock sync.Once

func InitFileHandlerRepo(r data_store.Datastore) {
	muLock.Do(func() {
		repo = r
	})
}
