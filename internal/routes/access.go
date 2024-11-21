package routes

import (
	accessHandler "betis-oprec/internal/handlers/access"

	"github.com/gofiber/fiber/v2"
)

// SetupAccessRoutes sets up the routes for the access handler
func SetupAccessRoutes(router fiber.Router) {
	access := router.Group("/access")

	access.Post("/", accessHandler.CreateAccess)
	access.Delete("/", accessHandler.DeleteAccess)
	access.Get("/", accessHandler.GetAccess)
}
