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

	delBook := model.MagicBook{
		ID: id,
	}

	if err := db.Delete(&delBook, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete book!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
