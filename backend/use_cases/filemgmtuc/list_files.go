package filemgmtuc

import (
	"backend/domain/entities"
	"context"
)

type getFilesUseCase struct {
	fileManager fileManager
}

// /test commit
func NewGetFilesUseCase(fm fileManager) *getFilesUseCase {
	return &getFilesUseCase{
		fileManager: fm,
	}
}

func (uc *getFilesUseCase) GetFiles(ctx context.Context, userId string) ([]entities.File, error) {

	files, err := uc.fileManager.List(ctx, userId)
	if err != nil {
		return nil, err
	}
	return files, err
}
