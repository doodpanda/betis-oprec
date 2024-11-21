package accessHandlers

import (
	"betis-oprec/database"
	"betis-oprec/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateAccess handles the creation of a new access permission for a witch to a magic book.
// @Summary Create a new access permission
// @Description Create a new access permission for a witch to a magic book
// @Tags access
// @Accept json
// @Produce json
// @Param access body AccessCreateRequest true "Access Create Request"
// @Success 200 {object} model.AccessPermission
// @Failure 400 {object} fiber.Map{"message": string}
// @Failure 404 {object} fiber.Map{"message": string}
// @Failure 500 {object} fiber.Map{"message": string}
// @Router /access [post]
func CreateAccess(c *fiber.Ctx) error {
	db := database.DB

	access := new(AccessCreateRequest)
	if err := c.BodyParser(access); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Malformed request",
		})
	}

	witchID, err := uuid.Parse(access.WitchID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid WitchID",
		})
	}

	magicBookID, err := uuid.Parse(access.MagicBookID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid MagicBookID",
		})
	}

	var witch model.Witch
	if err := db.First(&witch, "id = ?", witchID).Error; err != nil {
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
	if err := db.First(&magicBook, "id = ?", magicBookID).Error; err != nil {
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
		WitchID:     witchID,
		MagicBookID: magicBookID,
		PermitDate:  time.Now(),
	}

	if err := db.Create(&newAccess).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create access!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(newAccess)
}
