package filemgmtuc

import (
	"context"
)

type downloadFileUseCase struct {
	fileManager fileManager
}

func NewDownloadFileUseCase(fm fileManager) *downloadFileUseCase {
	return &downloadFileUseCase{
		fileManager: fm,
	}
}

func (uc *downloadFileUseCase) DownloadFile(ctx context.Context, userId string, filename string) (string, error) {

	files, err := uc.fileManager.Download(ctx, filename, userId)
	if err != nil {
		return "", err
	}
	return files, err
}
