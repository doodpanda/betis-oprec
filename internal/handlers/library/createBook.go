package libraryHandlers

import (
	"time"

	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BookCreate handles the creation of a new MagicBook.
func BookCreate(c *fiber.Ctx) error {
	db := database.DB

	// Parse request body into the BookCreateRequest struct
	book := new(BookCreateRequest)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	// Validate MagicType input
	if book.MagicType != "elemental" && book.MagicType != "illusion" && book.MagicType != "necromancy" && book.MagicType != "healing" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid magic type",
		})
	}

	// Create new MagicBook instance
	newBook := model.MagicBook{
		ID:          uuid.New(),
		Title:       book.Title,
		MagicType:   string(book.MagicType),
		CreatedDate: time.Now(),
		IsAvailable: book.Status,
	}

	// Save the book to the database
	if err := db.Create(&newBook).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add book!",
		})
	}

	// Return the created book as a response
	return c.Status(fiber.StatusCreated).JSON(newBook)
}
