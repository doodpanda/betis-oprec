package accessHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAccess(c *fiber.Ctx) error {
	db := database.DB

	access := new(AccessCreateRequest)
	if err := c.BodyParser(access); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	newAccess := model.AccessPermission{
		ID:          uuid.New(),
		WitchID:     uuid.MustParse(access.WitchID),
		MagicBookID: uuid.MustParse(access.MagicBookID),
		PermitDate:  time.Now(),
	}

	if err := db.Create(&newAccess).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create access!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(access)
}
