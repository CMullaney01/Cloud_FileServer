package routes

import (
	"backend/api/handlers"
	"backend/infrastructure/identity"
	"backend/infrastructure/mongomgmt"
	"backend/use_cases/filemgmtuc"
	"backend/use_cases/usermgmtuc"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to My File Server Rest API"))
	})

	grp := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	grp.Post("/user", handlers.RegisterHandler(registerUseCase))

}

func InitProtectedRoutes(app *fiber.App) {

	grp := app.Group("/api/v1")

	fileManager := mongomgmt.NewFileManager()

	uploadFileUseCase := filemgmtuc.NewFileUploadUseCase(fileManager)
	grp.Post("/files", handlers.UploadFileHandler(uploadFileUseCase))

	downloadFileUseCase := filemgmtuc.NewDownloadFileUseCase(fileManager)
	grp.Get("/files", handlers.DownloadFileHandler(downloadFileUseCase))

	getFilesUseCase := filemgmtuc.NewGetFilesUseCase(fileManager)
	grp.Get("/filelist", handlers.GetFilesHandler(getFilesUseCase))
}
