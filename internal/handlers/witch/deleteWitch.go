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

	delWitch := model.Witch{
		ID: id,
	}

	if err := db.Delete(&delWitch, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete witch!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
