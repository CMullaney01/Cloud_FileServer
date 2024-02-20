package filemgmtuc

import (
	"backend/domain/entities"
)

type FileUploadDownload interface {
	// List() []entities.File
	// Download() entities.File
	Upload(file *entities.File) error
}
