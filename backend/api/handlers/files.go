package handlers

import (
	"context"

	"backend/shared/enums"
	"backend/use_cases/filemgmtuc"

	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type UploadFileUseCase interface {
	UploadFile(ctx context.Context, request filemgmtuc.UploadFileRequest) (*filemgmtuc.UploadFileResponse, error)
}

type GetFilesUseCase interface {
	GetFiles(ctx context.Context, userId string) (map[string]string, error)
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

func GetFilesHandler(useCase GetFilesUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var ctx = c.UserContext()

		// Retrieve the claims from the context
		claims, ok := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Access the "username" field from the claims
		username, ok := claims["username"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		files, err := useCase.GetFiles(ctx, username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(files)
	}
}
