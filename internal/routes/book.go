package routes

import (
	libraryHandler "betis-oprec/internal/handlers/library"
	//witchHandler "betis-oprec/internal/handlers/witch" // Import the witch handler

	"github.com/gofiber/fiber/v2"
)

func SetupLibraryRoutes(router fiber.Router) {
	library := router.Group("/library")

	library.Post("/", libraryHandler.BookCreate)
	library.Delete("/", libraryHandler.DeleteBook)
	library.Get("/", libraryHandler.GetBook)
	library.Patch("/", libraryHandler.UpdateBook)
	//library.Post("/witch", witchHandler.WitchCreate) // Add route for creating a witch
}
