package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"errors"
	"fmt"
)

func UpsertToken(req models.Token) error {
	fmt.Println("check res ", req)
	err := db.DB.Unscoped().Table("tokens").Save(req).Error
	if err != nil {
		return errors.New("failed while upserting token")
	}
	return nil
}
