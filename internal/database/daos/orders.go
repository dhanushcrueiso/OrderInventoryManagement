package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"errors"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func UpsertOrders(req []models.Order) error {
	err := db.DB.Unscoped().Table("orders").Save(&req).Error
	if err != nil {
		return errors.New("unable to place order")
	}
	return nil
}

func GetAllOrdersWithCustomerId(c *fiber.Ctx, customerId uuid.UUID) ([]dtos.Order, error) {
	res := []dtos.Order{}
	q := db.DB.Unscoped().Table("orders").Where("customer_id = ?", customerId)
	if c.Query("order_id") != "" {
		q.Where("order_id = ?", c.Query("order_id"))
	}
	err := q.Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
