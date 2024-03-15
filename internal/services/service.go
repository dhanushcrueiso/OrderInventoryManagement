package services

import (
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/dtos"
	"fmt"

	"OrderInventoryManagement/internal/database/models"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func CheckName(c *fiber.Ctx, req string) error {

	res, err := daos.GetName(c, req)
	if err != nil {
		return err
	}
	fmt.Println("res", res)
	return nil
}

func SaveUser(c *fiber.Ctx, req dtos.User) error {
	daoReq, err := DtosToDao(req)
	if err != nil {
		return err
	}

	if err := daos.SaveUser(c, daoReq); err != nil {
		return err
	}
	return nil
}

func DtosToDao(req dtos.User) (models.User, error) {

	return models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
		Mobile:   req.Mobile,
	}, nil
}
