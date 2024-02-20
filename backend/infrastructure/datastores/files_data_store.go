package datastores

import (
	"backend/domain/entities"
	"sync"

	"github.com/google/uuid"
)

type filesDataStore struct {
	files map[uuid.UUID]entities.File
	sync.Mutex
}

func NewFilesDataStore() *filesDataStore {
	return &filesDataStore{
		files: make(map[uuid.UUID]entities.File),
	}
}

// takes instance of the product and saves it in a map in memory
func (ds *filesDataStore) Upload(file *entities.File) error {
	ds.Lock()
	ds.files[file.Id] = *file
	ds.Unlock()
	return nil
}
