package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"errors"

	"github.com/gofiber/fiber"
)

func SaveUser(c *fiber.Ctx, req models.User) error {

	err := db.DB.Debug().Unscoped().Table("users").Save(req).Error
	if err != nil {
		return err
	}

	return err
}

func GetAccount(c *fiber.Ctx, req dtos.User) (*models.User, error) {
	user := &models.User{}
	err := db.DB.Debug().Unscoped().Table("users").Where("username = ?", req.Username).First(user).Error
	if err != nil {
		return nil, errors.New("unable to fetch users")
	}

	return user, nil
}

func GetAccountById(c *fiber.Ctx, req dtos.User) (*models.User, error) {
	user := &models.User{}
	err := db.DB.Debug().Unscoped().Table("users").Where("id = ?", req.ID).First(user).Error
	if err != nil {
		return nil, errors.New("unable to fetch users")
	}

	return user, nil
}
