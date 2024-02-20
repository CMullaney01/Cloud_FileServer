package filemgmtuc

import (
	"backend/domain/entities"
)

type FileUploadDownload interface {
	ListFiles() []entities.File
	DownloadFile() entities.File
	UploadFile(file *entities.File) error
}
