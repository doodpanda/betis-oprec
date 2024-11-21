package routes

import (
	witchHandler "betis-oprec/internal/handlers/witch" // Import the witch handler

	"github.com/gofiber/fiber/v2"
)

// SetupWitchRoutes sets up the routes for the witch handler
func SetupWitchRoutes(router fiber.Router) {
	witch := router.Group("/witch")

	witch.Post("/", witchHandler.CreateWitch)
	witch.Delete("/", witchHandler.DeleteWitch)
	witch.Get("/", witchHandler.GetWitch)
	witch.Patch("/", witchHandler.UpdateWitch)
}
