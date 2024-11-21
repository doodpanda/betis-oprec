package witchHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateWitch updates an existing witch's details in the database.
//
// @Summary Update witch
// @Description Update the details of an existing witch by ID.
// @Tags witches
// @Accept json
// @Produce json
// @Param witch body WitchUpdateRequest true "Witch Update Request"
// @Success 200 {object} model.Witch "Successfully updated witch"
// @Failure 400 {object} fiber.Map "Malformed request or invalid UUID format"
// @Failure 404 {object} fiber.Map "Witch not found"
// @Failure 500 {object} fiber.Map "Failed to modify witch"
// @Router /witch/update [put]
func UpdateWitch(c *fiber.Ctx) error {
	db := database.DB

	witch := new(WitchUpdateRequest)
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

	existingWitch.Name = witch.Name
	existingWitch.Age = witch.Age
	existingWitch.Rank = witch.Rank

	if err := db.Save(&existingWitch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to modify witch!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(existingWitch)
}
