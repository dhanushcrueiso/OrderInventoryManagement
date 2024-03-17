package handlers

import (
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func GetProducts(c *fiber.Ctx) {
	res, err := services.GetAllProductsList(c)
	if err != nil {
		return
	}
	c.JSON(res)
}

func AddProducts(c *fiber.Ctx) {
	var req dtos.Product
	if err := c.BodyParser(&req); err != nil {
		return
	}
	fmt.Println("checking the parser", req)
	//checking if the user is admin to allow him to add products
	aid, _ := uuid.Parse(c.Params("aid"))
	admin, err := daos.GetAccountById(c, dtos.User{ID: aid})
	if err != nil {
		c.JSON(dtos.Response{Code: fiber.StatusNotFound, Message: "user not an admin"})
	}
	if admin.Role != "admin" {
		c.JSON(dtos.Response{Code: fiber.StatusNotFound, Message: "user not an admin 1"})
	}
	err = services.AddProducts(c, req)
	if err != nil {
		return
	}
	c.JSON(dtos.Response{Code: fiber.StatusOK, Message: "product udpated"})
}
