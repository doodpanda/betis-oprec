package router

import (
	routes "betis-oprec/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	routes.SetupLibraryRoutes(api)
	routes.SetupWitchRoutes(api)
	routes.SetupAccessRoutes(api)
}
