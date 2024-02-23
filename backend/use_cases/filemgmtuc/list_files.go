package filemgmtuc

import (
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

func (uc *getFilesUseCase) GetFiles(ctx context.Context, userId string) (map[string]string, error) {

	files, err := uc.fileManager.List(ctx, userId)
	if err != nil {
		return nil, err
	}
	return files, err
}
