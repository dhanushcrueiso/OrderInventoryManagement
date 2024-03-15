package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"

	"github.com/gofiber/fiber"
)

func SaveUser(c *fiber.Ctx, req models.User) error {

	err := db.DB.Unscoped().Table("user").Save(req).Error
	if err != nil {
		return err
	}

	return err
}
