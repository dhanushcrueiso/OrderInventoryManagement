package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAll(c *fiber.Ctx) ([]*models.Product, error) {
	products := []*models.Product{}
	err := db.DB.Unscoped().Debug().Table("products p").Select("p.id,p.name,p.description,p.price,i.quantity").Joins("JOIN inventory i on p.id = i.product_id").Scan(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func Upsert(c *fiber.Ctx, req models.Product) (uuid.UUID, error) {
	check := models.Product{}
	var pid = uuid.Nil
	updateprod := models.ProductUpdate{}
	err := db.DB.Unscoped().Table("products").Where("id =?", req.ID).First(&check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return pid, err
	}
	if err == nil {

		updateprod.ID = check.ID
		pid = updateprod.ID
		updateprod.Description = req.Description
		updateprod.Price = req.Price
		updateprod.Name = req.Name
		err := db.DB.Unscoped().Table("products").Where("id =?", updateprod.ID).Save(&updateprod).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return pid, err
		}
		return pid, nil

	}
	updateprod.ID = uuid.New()
	updateprod.Description = req.Description
	updateprod.Name = req.Name
	updateprod.Price = req.Price
	pid = updateprod.ID
	err = db.DB.Unscoped().Debug().Table("products").Create(updateprod).Error
	if err != nil {
		return pid, err
	}
	return pid, nil
}

func UpdateInventory(c *fiber.Ctx, req models.Inventory, update string) error {
	check := models.Inventory{}
	err := db.DB.Unscoped().Table("inventory").Where("product_id = ?", req.ProductID).First(&check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil && update != "subtract" {

		req.ID = check.ID
		req.LastUpdated = time.Now()
		fmt.Println("inside err nil", req)
		err := db.DB.Unscoped().Table("inventory").Save(&req).Error
		if err != nil {
			return err
		}
		return nil
	}
	if err == nil && update == "subtract" {
		fmt.Println("inside check ", check)
		fmt.Println("inside req", req)

		c := check.Quantity - req.Quantity
		subreq := models.Inventory{
			ID:          check.ID,
			ProductID:   req.ProductID,
			Quantity:    c,
			LastUpdated: time.Now(),
		}
		fmt.Println("inside", c)
		fmt.Println("inside err nil", subreq)
		err := db.DB.Unscoped().Table("inventory").Save(&subreq).Error
		if err != nil {
			return err
		}
		return nil
	}

	req.ID = uuid.New()
	req.LastUpdated = time.Now()
	fmt.Println("create place", req)
	err = db.DB.Unscoped().Debug().Table("inventory").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}
