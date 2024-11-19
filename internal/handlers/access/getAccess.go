package accessHandlers

import (
	"betis-oprec/database"

	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetAccess(c *fiber.Ctx) error {
	db := database.DB

	accessID := c.Query("id")
	witchID := c.Query("witchId")
	bookID := c.Query("bookId")

	if witchID != "" {
		var accesses []model.AccessPermission
		if err := db.Find(&accesses, "witchId = ?", witchID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}
		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	if bookID != "" {
		var accesses []model.AccessPermission
		if err := db.Find(&accesses, "bookId = ?", bookID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}
		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	if accessID == "" {
		var accesses []model.AccessPermission
		if err := db.Find(&accesses).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}
		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	var access model.AccessPermission
	if err := db.First(&access, "id = ?", accessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Access not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(access)
}
