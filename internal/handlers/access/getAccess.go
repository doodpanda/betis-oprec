package accessHandlers

import (
	"betis-oprec/database"

	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

// GetAccess handles the retrieval of access permissions based on query parameters.
// @Summary Retrieve access permissions
// @Description Retrieves access permissions based on provided query parameters such as id, witch_id, and magic_book_id. Supports pagination.
// @Tags access
// @Accept json
// @Produce json
// @Param id query string false "Access ID"
// @Param witch_id query string false "Witch ID"
// @Param magic_book_id query string false "Magic Book ID"
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {object} AccessListResponse "List of access permissions"
// @Success 200 {object} AccessInfoResponse "Single access permission"
// @Failure 404 {object} fiber.Map "Failed to retrieve access"
// @Router /access [get]
func GetAccess(c *fiber.Ctx) error {
	db := database.DB

	accessID := c.Query("id")
	witchID := c.Query("witch_id")
	bookID := c.Query("magic_book_id")

	if bookID != "" || witchID != "" {
		limit := c.QueryInt("limit", 10) // default limit is 10
		page := c.QueryInt("page", 1)    // default page is 1
		offset := (page - 1) * limit

		var accesses AccessListResponse
		if err := db.Limit(limit).Offset(offset).Find(&accesses.Accesses, "magic_book_id = ? OR witch_id = ?", bookID, witchID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}
		var total int64
		db.Model(&model.AccessPermission{}).Where("magic_book_id = ? OR witch_id = ?", bookID, witchID).Count(&total)
		accesses.Total = int(total)
		accesses.HasNext = accesses.Total*page < int(total)
		accesses.Page = page
		accesses.TotalPage = int(total)/limit + 1

		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	if accessID == "" {
		limit := c.QueryInt("limit", 10) // default limit is 10
		page := c.QueryInt("page", 1)    // default page is 1
		offset := (page - 1) * limit

		var accesses AccessListResponse
		if err := db.Limit(limit).Offset(offset).Find(&accesses.Accesses).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}

		var total int64
		db.Model(&model.AccessPermission{}).Count(&total)
		accesses.Total = int(total)
		accesses.HasNext = accesses.Total*page < int(total)
		accesses.Page = page
		accesses.TotalPage = int(total)/limit + 1
		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	var access AccessInfoResponse
	var accessSingle model.AccessPermission
	if err := db.First(&accessSingle, "id = ?", accessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Access not found",
		})
	}
	access.ID = accessSingle.ID
	access.WitchID = accessSingle.WitchID
	access.MagicBookID = accessSingle.MagicBookID
	access.PermitDate = accessSingle.PermitDate
	if err := db.First(&access.Witch, "id = ?", access.WitchID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to get detail witch!",
		})
	}
	if err := db.First(&access.MagicBook, "id = ?", access.MagicBookID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to get detail book!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(access)
}
