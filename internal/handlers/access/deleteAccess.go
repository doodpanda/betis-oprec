package accessHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

func DeleteAccess(c *fiber.Ctx) error {
	db := database.DB

	accessID := c.Query("id")

	var access model.AccessPermission
	if err := db.First(&access, "id = ?", accessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Access not found",
		})
	}

	if err := db.Delete(&access).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete access!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Access deleted successfully",
	})
}
