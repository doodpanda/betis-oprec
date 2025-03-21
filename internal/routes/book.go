package routes

import (
	libraryHandler "betis-oprec/internal/handlers/library"

	"github.com/gofiber/fiber/v2"
)

// SetupLibraryRoutes sets up the routes for the book handler
func SetupLibraryRoutes(router fiber.Router) {
	library := router.Group("/library")

	library.Post("/", libraryHandler.CreateBook)
	library.Delete("/", libraryHandler.DeleteBook)
	library.Get("/", libraryHandler.GetBook)
	library.Patch("/", libraryHandler.UpdateBook)
}
