package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"errors"
)

func UpsertOrders(req []models.Order) error {
	err := db.DB.Unscoped().Table("orders").Save(&req).Error
	if err != nil {
		return errors.New("unable to place order")
	}
	return nil
}
