package accessHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

// DeleteAccess handles the deletion of an access permission.
// @Summary Delete an access permission
// @Description Deletes an access permission based on the provided access ID.
// @Tags access
// @Accept json
// @Produce json
// @Param id query string true "Access ID"
// @Success 200 {object} map[string]interface{} "Access deleted successfully"
// @Failure 400 {object} map[string]interface{} "Access ID is required"
// @Failure 404 {object} map[string]interface{} "Access not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete access"
// @Router /access [delete]
func DeleteAccess(c *fiber.Ctx) error {
	db := database.DB

	accessID := c.Query("id")
	if accessID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Access ID is required",
		})
	}

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
