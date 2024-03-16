package handlers

import (
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

func CustomerSignup(c *fiber.Ctx) {
	fmt.Println("here:")
	var customer dtos.Customer
	if err := c.BodyParser(&customer); err != nil {
		return
	}
	fmt.Println("checking the parser", customer)

	if err := services.SaveCustomer(c, customer); err != nil {
		return
	}
	c.JSON(dtos.Response{Code: http.StatusOK, Message: "signup successfull"})
}

func CustomerLogin(c *fiber.Ctx) {
	var customer dtos.Customer
	if err := c.BodyParser(&customer); err != nil {
		return
	}
	fmt.Println("checking the parser", customer)

	res, err := services.CustomerLogin(c, customer)
	if err != nil {
		return
	}
	c.JSON(dtos.Login{AccountInfo: *res, Code: http.StatusOK, Message: "login successfull"})
}
