package witchHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteWitch(c *fiber.Ctx) error {
	db := database.DB

	witch := new(WitchDeleteRequest)
	if err := c.BodyParser(witch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	id, err := uuid.Parse(witch.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID format",
		})
	}

	var existingWitch model.Witch
	if err := db.First(&existingWitch, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Witch not found",
		})
	}

	if err := db.Delete(&existingWitch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete witch!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
