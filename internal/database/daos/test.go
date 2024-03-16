package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"log"

	"github.com/gofiber/fiber"
)

func GetName(c *fiber.Ctx, req string) (models.Accountcheck, error) {
	res := models.Accountcheck{}
	err := db.DB.Unscoped().Table("account").Where("id = ?", req).Take(&res).Error
	if err != nil {
		log.Fatalln("unable to fetch data ")
		return res, err
	}
	return res, nil
}
