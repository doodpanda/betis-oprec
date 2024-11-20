package libraryHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
