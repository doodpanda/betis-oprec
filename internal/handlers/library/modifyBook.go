package libraryHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateBook updates an existing book in the database.
// @Summary Update an existing book
// @Description Update the details of an existing book by its UUID
// @Tags books
// @Accept json
// @Produce json
// @Param book body BookUpdateRequest true "Book Update Request"
// @Success 200 {object} model.MagicBook "Successfully updated book"
// @Failure 400 {object} fiber.Map "Malformed request or invalid UUID format"
// @Failure 404 {object} fiber.Map "Book not found"
// @Failure 500 {object} fiber.Map "Failed to modify book"
// @Router /books [put]
func UpdateBook(c *fiber.Ctx) error {
	db := database.DB

	book := new(BookUpdateRequest)
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

	existingBook.Title = book.Title
	existingBook.MagicType = string(book.MagicType)
	existingBook.IsAvailable = book.Status

	if err := db.Save(&existingBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to modify book!",
		})
	}

	modBook := existingBook

	return c.Status(fiber.StatusOK).JSON(modBook)
}
