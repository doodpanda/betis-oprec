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
	if err := c.BodyParser(access); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	var witch model.Witch
	if err := db.First(&witch, "id = ?", uuid.MustParse(access.WitchID)).Error; err != nil {
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
	if err := db.First(&magicBook, "id = ?", uuid.MustParse(access.MagicBookID)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Magic book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to query magic book",
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

	return c.Status(fiber.StatusOK).JSON(newAccess)
}
