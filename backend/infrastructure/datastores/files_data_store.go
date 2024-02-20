package datastores

import (
	"backend/domain/entities"
	"sort"
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
func (ds *filesDataStore) Create(file *entities.File) error {
	ds.Lock()
	ds.files[file.Id] = *file
	ds.Unlock()
	return nil
}

func (ds *filesDataStore) GetAll() []entities.File {
	all := make([]entities.File, 0, len(ds.files))
	for _, value := range ds.files {
		all = append(all, value)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt.Before(all[j].CreatedAt)
	})
	return all
}
