package routes

import (
	"backend/api/handlers"
	"backend/api/middlewares"
	"backend/infrastructure/datastores"
	"backend/infrastructure/identity"
	"backend/infrastructure/mongomgmt"
	"backend/use_cases/filemgmtuc"
	"backend/use_cases/productsuc"
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

	fileManager := mongomgmt.NewFileManager()
	// auths for if we didnt need role-based auth
	uploadFileUseCase := filemgmtuc.NewFileUploadUseCase(fileManager)
	grp.Post("/files", handlers.UploadFileHandler(uploadFileUseCase))
}

func InitProtectedRoutes(app *fiber.App) {

	grp := app.Group("/api/v1")

	productsDataStore := datastores.NewProductsDataStore()

	createProductUseCase := productsuc.NewCreateProductUseCase(productsDataStore)
	grp.Post("/products", middlewares.NewRequiresRealmRole("admin"),
		handlers.CreateProductHandler(createProductUseCase))

	getProductsUseCase := productsuc.NewGetProductsUseCase(productsDataStore)
	grp.Get("/products", middlewares.NewRequiresRealmRole("viewer"),
		handlers.GetProductsHandler(getProductsUseCase))

}
