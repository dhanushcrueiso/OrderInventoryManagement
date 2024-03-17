package services

import (
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func GetAllProductsList(c *fiber.Ctx) ([]*dtos.Product, error) {
	res, err := daos.GetAll(c)
	if err != nil {
		return nil, err
	}
	products, err := ProductDaotoDtos(res)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func AddProducts(c *fiber.Ctx, req dtos.Product) error {
	daoReq, _ := ProductDtosToDaos(req)
	pid, err := daos.Upsert(c, daoReq)
	if err != nil {
		return err
	}

	err = daos.UpdateInventory(c, models.Inventory{
		ID:          uuid.New(),
		ProductID:   pid,
		Quantity:    req.Quantity,
		LastUpdated: time.Now(),
	}, "add")
	if err != nil {
		return err
	}
	return nil

}

func ProductDaotoDtos(req []*models.Product) ([]*dtos.Product, error) {
	res := []*dtos.Product{}
	for _, product := range req {
		pro := dtos.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		}
		res = append(res, &pro)

	}
	return res, nil
}

func ProductDtosToDaos(req dtos.Product) (models.Product, error) {

	pro := models.Product{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	return pro, nil
}
