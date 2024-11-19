package routes

import (
	accessHandler "betis-oprec/internal/handlers/access"

	"github.com/gofiber/fiber/v2"
)

// SetupAccessRoutes sets up the routes for the access handler
func SetupAccessRoutes(router fiber.Router) {
	access := router.Group("/access") // Create a new group for the access handler

	access.Post("/", accessHandler.CreateAccess)   // Create a new access
	access.Delete("/", accessHandler.DeleteAccess) // Delete an access
	access.Get("/", accessHandler.GetAccess)       // Get an access
}
