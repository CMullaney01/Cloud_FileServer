package handlers

import (
	"context"

	"backend/use_cases/filemgmtuc"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type UploadFileUseCase interface {
	UploadFile(ctx context.Context, request filemgmtuc.UploadFileRequest) (*filemgmtuc.UploadFileResponse, error)
}

func UploadFileHandler(useCase UploadFileUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		var request = filemgmtuc.UploadFileRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.UploadFile(ctx, request)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}
