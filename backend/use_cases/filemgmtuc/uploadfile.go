package filemgmtuc

import (
	"backend/domain/entities"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UploadFileRequest struct {
	Name string `validate:"required,min=3,max=15"`
	Path string `validate:"required"`
}

type UploadFileResponse struct {
	File *entities.File
}

type uploadFileUseCase struct {
	fileStore FileUploadDownload
}

func NewFileUploadUseCase(fs FileUploadDownload) *uploadFileUseCase {
	return &uploadFileUseCase{
		fileStore: fs,
	}
}

func (uc *uploadFileUseCase) UploadFile(ctx context.Context, request UploadFileRequest) (*UploadFileResponse, error) {

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var file = &entities.File{
		Id:          uuid.New(),
		CreatedAt:   time.Now(),
		Name:        request.Name,
		ContentType: "pdf",
		Size:        0,
		Path:        request.Path,
	}

	err = uc.fileStore.UploadFile(file)
	if err != nil {
		return nil, err
	}

	var response = &UploadFileResponse{File: file}
	return response, nil
}
