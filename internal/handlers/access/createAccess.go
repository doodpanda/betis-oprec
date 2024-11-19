package accessHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateAccess(c *fiber.Ctx) error {
	db := database.DB

	access := new(AccessCreateRequest)

	var witch model.Witch
	if err := db.First(&witch, "id = ?", access.WitchID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Witch not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to query witch",
		})
	}

	var magicBook model.MagicBook
	if err := db.First(&magicBook, "id = ?", access.MagicBookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Magic book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to query magic book",
		})
	}
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
