package witchHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetWitch(c *fiber.Ctx) error {
	db := database.DB

	witchID := c.Query("id")

	if witchID == "" {
		var witches []model.Witch
		if err := db.Find(&witches).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve books",
			})
		}
		return c.Status(fiber.StatusOK).JSON(witches)
	}

	var witch model.Witch
	if err := db.First(&witch, "id = ?", witchID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Witch not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(witch)
}
