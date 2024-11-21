package libraryHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteBook handles the deletion of a book from the library.
// @Summary Delete a book
// @Description Deletes a book from the library based on the provided UUID.
// @Tags books
// @Accept json
// @Produce json
// @Param book body BookDeleteRequest true "Book Delete Request"
// @Success 200 {object} map[string]interface{} "id"
// @Failure 400 {object} map[string]interface{} "Malformed request or Invalid UUID format"
// @Failure 404 {object} map[string]interface{} "Book not found"
// @Failure 500 {object} map[string]interface{} "Failed to delete book"
// @Router /books [delete]

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB

	book := new(BookDeleteRequest)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	id, err := uuid.Parse(book.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID format",
		})
	}

	var existingBook model.MagicBook
	if err := db.First(&existingBook, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	if err := db.Delete(&existingBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete book!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
