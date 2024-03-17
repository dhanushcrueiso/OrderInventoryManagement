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
	//fetching products along with available quantity from inventory
	q := db.DB.Unscoped().Debug().Table("products p").Select("p.id,p.name,p.description,p.price,i.quantity").Joins("JOIN inventory i on p.id = i.product_id")
	if c.Query("q") != "" {
		q.Where("p.name ilike ?", c.Query("q")+"%")
	}

	err := q.Scan(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func Upsert(c *fiber.Ctx, req models.Product) (uuid.UUID, error) {
	check := models.Product{}
	var pid = uuid.Nil
	updateprod := models.ProductUpdate{}
	//checking if the product with id already exists
	err := db.DB.Unscoped().Table("products").Where("id =?", req.ID).First(&check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return pid, err
	}
	if err == nil {
		//if exists then we update the data
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
	//if it does not exists we create a new row of data
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
		//checking if the update == subtract which basically means order is being placed and inventory item needs to be subtracted
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
