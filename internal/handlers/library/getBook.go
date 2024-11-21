package libraryHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

// GetBook handles the retrieval of a book or a list of books from the database.
//
// @Summary Retrieve a book or a list of books
// @Description This endpoint retrieves a single book by its ID or a paginated list of books if no ID is provided.
// @Tags books
// @Accept json
// @Produce json
// @Param id query string false "Book ID"
// @Param limit query int false "Number of books to retrieve per page" default(10)
// @Param page query int false "Page number to retrieve" default(1)
// @Success 200 {object} model.MagicBook "Single book retrieved successfully"
// @Success 200 {object} BookListResponse "List of books retrieved successfully"
// @Failure 404 {object} fiber.Map "Book not found or failed to retrieve books"
// @Router /books [get]

func GetBook(c *fiber.Ctx) error {
	db := database.DB

	// Get the book ID from the query parameters
	bookID := c.Query("id")

	// Validate the book ID
	if bookID == "" {
		limit := c.QueryInt("limit", 10) // default limit is 10
		page := c.QueryInt("page", 1)    // default page is 1
		offset := (page - 1) * limit

		var magicBooks BookListResponse
		if err := db.Limit(limit).Offset(offset).Find(&magicBooks.Books).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve books",
			})
		}

		// Get the total number of books and determine if there are more books
		var total int64
		db.Model(&model.MagicBook{}).Count(&total)              // Gets the total count of books in the database
		magicBooks.Total = int(total)                           // Number of books fetched in the current operation
		magicBooks.HasNext = magicBooks.Total*page < int(total) // Check if there are more books to fetch
		magicBooks.Page = page                                  // Current page
		magicBooks.TotalPage = int(total)/limit + 1             // Total number of pages

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
