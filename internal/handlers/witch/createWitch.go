package witchHandlers

import (
	"betis-oprec/database"

	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// WitchCreate handles the creation of a new Witch.
// CreateWitch handles the creation of a new witch.
//
// @Summary Create a new witch
// @Description This endpoint creates a new witch with the provided details in the request body.
// @Tags witches
// @Accept json
// @Produce json
// @Param witch body WitchCreateRequest true "Witch Create Request"
// @Success 201 {object} model.Witch "Created witch"
// @Failure 400 {object} fiber.Map "Malformed request or invalid role"
// @Failure 500 {object} fiber.Map "Failed to add witch"
// @Router /witches [post]
func CreateWitch(c *fiber.Ctx) error {
	db := database.DB

	// Parse request body into the WitchCreateRequest struct
	witch := new(WitchCreateRequest)
	if err := c.BodyParser(witch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	// Validate Role input
	if witch.Rank != "apprentice" && witch.Rank != "adept" && witch.Rank != "master" && witch.Rank != "archmage" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid role",
		})
	}

	// Create new Witch instance
	newWitch := model.Witch{
		ID:   uuid.New(),
		Name: witch.Name,
		Rank: witch.Rank,
		Age:  witch.Age,
	}

	// Save the witch to the database
	if err := db.Create(&newWitch).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add witch!",
		})
	}

	// Return the created witch as a response
	return c.Status(fiber.StatusCreated).JSON(newWitch)
}
