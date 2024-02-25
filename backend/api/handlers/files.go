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
	UploadFile(ctx context.Context, userId string, request filemgmtuc.UploadFileRequest) (*filemgmtuc.UploadFileResponse, error)
}

type DownloadFileUseCase interface {
	DownloadFile(ctx context.Context, userId string, fileId string) (string, error)
}

type GetFilesUseCase interface {
	GetFiles(ctx context.Context, userId string) (map[string]string, error)
}

func UploadFileHandler(useCase UploadFileUseCase) fiber.Handler {
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
		username, ok := claims["preferred_username"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		var request = filemgmtuc.UploadFileRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.UploadFile(ctx, username, request)
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
		username, ok := claims["preferred_username"].(string)
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

func DownloadFileHandler(useCase DownloadFileUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Access the query parameter "filename"
		filename := c.Query("filename")

		// Check if filename is empty or not provided
		if filename == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Filename query parameter is required",
			})
		}

		var ctx = c.UserContext()

		// Retrieve the claims from the context
		claims, ok := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Access the "username" field from the claims
		username, ok := claims["preferred_username"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		files, err := useCase.DownloadFile(ctx, username, filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(files)
	}
}
