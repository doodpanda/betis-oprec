package accessHandlers

import (
	"betis-oprec/database"

	"betis-oprec/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetAccess(c *fiber.Ctx) error {
	db := database.DB

	accessID := c.Query("id")
	witchID := c.Query("witchId")
	bookID := c.Query("magicBookId")

	if witchID != "" {
		limit := c.QueryInt("limit", 10) // default limit is 10
		page := c.QueryInt("page", 1)    // default page is 1
		offset := (page - 1) * limit

		var accesses AccessListResponse
		if err := db.Limit(limit).Offset(offset).Find(&accesses.Accesses, "witch_id = ?", witchID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}

		accesses.Total = len(accesses.Accesses)
		accesses.HasNext = accesses.Total > limit

		return c.Status(fiber.StatusOK).JSON(accesses)
	}

	if bookID != "" {
		limit := c.QueryInt("limit", 10) // default limit is 10
		page := c.QueryInt("page", 1)    // default page is 1
		offset := (page - 1) * limit

		var accesses AccessListResponse
		if err := db.Limit(limit).Offset(offset).Find(&accesses.Accesses, "magic_book_id = ?", bookID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Failed to retrieve access",
			})
		}

		accesses.Total = len(accesses.Accesses)
		accesses.HasNext = accesses.Total > limit

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
	if err := db.First(&access, "id = ?", accessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Access not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(access)
}
