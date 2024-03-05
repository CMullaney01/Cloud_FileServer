package filemgmtuc

import (
	"backend/domain/entities"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadFileRequest struct {
	FileName    string `validate:"required"`
	ContentType string `validate:"required"`
	Size        int64  `validate:"required"`
}

type UploadFileResponse struct {
	File         *entities.File `json:"file"`
	PresignedURL string         `json:"presignedURL"`
}

type uploadFileUseCase struct {
	fileManager fileManager
}

func NewFileUploadUseCase(fm fileManager) *uploadFileUseCase {
	return &uploadFileUseCase{
		fileManager: fm,
	}
}

func (uc *uploadFileUseCase) UploadFile(ctx context.Context, userId string, request UploadFileRequest) (*UploadFileResponse, error) {

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	fileID := primitive.NewObjectID()
	createdAt := time.Now()
	isPublic := false // You can set this as required

	// Construct S3 object key with userID and filename
	objectKey := userId + "/" + request.FileName

	s3Bucket := viper.GetString("AWS_S3_BUCKET")

	var file = &entities.File{
		ID:          fileID,
		UserID:      userId,
		FileName:    request.FileName,
		S3Bucket:    s3Bucket,
		S3ObjectKey: objectKey,
		CreatedAt:   createdAt,
		IsPublic:    isPublic,
		Size:        request.Size,
		ContentType: request.ContentType,
	}

	presignedURL, err := uc.fileManager.Upload(ctx, file)
	if err != nil {
		return nil, err
	}

	var response = &UploadFileResponse{
		File:         file,
		PresignedURL: presignedURL,
	}
	return response, nil
}
func (uc *uploadFileUseCase) ConfirmFileUpload(ctx context.Context, userId string, request UploadFileRequest) error {
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return err
	}

	fileID := primitive.NewObjectID()
	createdAt := time.Now()
	isPublic := false // You can set this as required

	// Construct S3 object key with userID and filename
	objectKey := userId + "/" + request.FileName

	s3Bucket := viper.GetString("AWS_S3_BUCKET")

	var file = &entities.File{
		ID:          fileID,
		UserID:      userId,
		FileName:    request.FileName,
		S3Bucket:    s3Bucket,
		S3ObjectKey: objectKey,
		CreatedAt:   createdAt,
		IsPublic:    isPublic,
		Size:        request.Size,
		ContentType: request.ContentType,
	}

	err = uc.fileManager.ConfirmUpload(ctx, file)
	if err != nil {
		return err
	}

	return nil
}
