package filemgmtuc

import (
	"backend/domain/entities"
	"context"
)

type fileManager interface {
	List(ctx context.Context, userID string) (map[string]string, error)
	Upload(ctx context.Context, file *entities.File) (string, error)
	Download(ctx context.Context, userID string, filename string) (string, error)
}
