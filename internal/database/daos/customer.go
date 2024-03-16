package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"errors"

	"github.com/gofiber/fiber"
)

func SaveCustomer(c *fiber.Ctx, req models.Customer) error {

	err := db.DB.Debug().Unscoped().Table("customers").Save(req).Error
	if err != nil {
		return err
	}

	return err
}

func GetAccountCustomer(c *fiber.Ctx, req dtos.Customer) (*models.Customer, error) {
	customer := &models.Customer{}
	err := db.DB.Debug().Unscoped().Table("customers").Where("username = ?", req.Username).First(customer).Error
	if err != nil {
		return nil, errors.New("unable to fetch users")
	}

	return customer, nil

}
