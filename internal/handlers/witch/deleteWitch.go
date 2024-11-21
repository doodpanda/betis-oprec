package witchHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteWitch handles the deletion of a witch record from the database.
//
// @Summary Delete a witch
// @Description Deletes a witch record based on the provided UUID.
// @Tags witches
// @Accept json
// @Produce json
// @Param witch body WitchDeleteRequest true "Witch Delete Request"
// @Success 200 {object} map[string]interface{} "Successfully deleted witch"
// @Failure 400 {object} map[string]interface{} "Malformed request or invalid UUID format"
// @Failure 404 {object} map[string]interface{} "Witch not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete witch"
// @Router /witches/{id} [delete]
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
