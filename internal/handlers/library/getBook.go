package libraryHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

// getBook handles the retrieval of a MagicBook by its ID.
func GetBook(c *fiber.Ctx) error {
	db := database.DB

	// Get the book ID from the query parameters
	bookID := c.Query("id")

	// Validate the book ID
	if bookID == "" {
		var magicBooks []model.MagicBook
		if err := db.Find(&magicBooks).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve books",
			})
		}
		return c.Status(fiber.StatusOK).JSON(magicBooks)
	}

	// Retrieve the book from the database
	var magicBook model.MagicBook
	if err := db.First(&magicBook, "id = ?", bookID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// Return the retrieved book as a response
	return c.Status(fiber.StatusOK).JSON(magicBook)
}
