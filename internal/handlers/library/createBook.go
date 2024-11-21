package libraryHandlers

import (
	"time"

	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateBook handles the creation of a new magic book.
// @Summary Create a new magic book
// @Description Parses the request body to create a new magic book and saves it to the database.
// @Tags books
// @Accept json
// @Produce json
// @Param book body BookCreateRequest true "Book Create Request"
// @Success 201 {object} model.MagicBook
// @Failure 400 {object} fiber.Map "Malformed request or invalid magic type"
// @Failure 500 {object} fiber.Map "Failed to add book"
// @Router /books [post]

func CreateBook(c *fiber.Ctx) error {
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
