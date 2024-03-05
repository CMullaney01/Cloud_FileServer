package filemgmtuc

import (
	"backend/domain/entities"
	"context"
)

type fileManager interface {
	List(ctx context.Context, userID string) ([]entities.File, error)
	Upload(ctx context.Context, file *entities.File) (string, error)
	ConfirmUpload(ctx context.Context, file *entities.File) error
	Download(ctx context.Context, userID string, filename string) (string, error)
}
