package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"errors"
	"fmt"

	"github.com/gofiber/fiber"
	"go.uber.org/zap"
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

func ProductAnalytics(c *fiber.Ctx) ([]models.ProductQuantity, error) {
	//joining order with products to fetch the  most ordered top 10 products
	res := []models.ProductQuantity{}
	err := db.DB.Unscoped().Debug().Table("orders").Select("products.id, products.name, SUM(orders.quantity_ordered) AS total_quantity_ordered").
		Joins("JOIN products ON orders.product_id = products.id").
		Group("products.id, products.name").
		Order("total_quantity_ordered DESC").Limit(10).Find(&res).Error
	if err != nil {
		return res, err
	}
	fmt.Println("checking products", zap.Any("", res))
	return res, nil
}
