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

func (uc *downloadFileUseCase) DownloadFile(ctx context.Context, userId string, fileId string) (string, error) {

	files, err := uc.fileManager.Download(ctx, fileId, userId)
	if err != nil {
		return "", err
	}
	return files, err
}
