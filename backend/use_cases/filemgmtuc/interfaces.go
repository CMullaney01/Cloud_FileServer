package filemgmtuc

import (
	"backend/domain/entities"
	"context"
)

type fileManager interface {
	// List() []entities.File
	// Download() entities.File
	Upload(ctx context.Context, file *entities.File) (string, error)
}
