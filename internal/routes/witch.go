package routes

import (
	witchHandler "betis-oprec/internal/handlers/witch" // Import the witch handler

	"github.com/gofiber/fiber/v2"
)

// SetupWitchRoutes sets up the routes for the witch handler
func SetupWitchRoutes(router fiber.Router) {
	witch := router.Group("/witch") // Create a new group for the witch handler

	witch.Post("/", witchHandler.CreateWitch)   // Create a new witch
	witch.Delete("/", witchHandler.DeleteWitch) // Delete a witch
	witch.Get("/", witchHandler.GetWitch)       // Get a witch
	witch.Patch("/", witchHandler.UpdateWitch)  // Update a witch
}
