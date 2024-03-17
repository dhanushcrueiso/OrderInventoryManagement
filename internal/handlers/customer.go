package handlers

import (
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
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

func PlaceOrder(c *fiber.Ctx) {
	var products []dtos.Product

	if err := c.BodyParser(&products); err != nil {
		return
	}
	cid, _ := uuid.Parse(c.Params("cid"))
	res, err := services.PlaceOrder(c, products, cid)
	if err != nil {
		return
	}
	mes := fmt.Sprintf("order places with order id %s", res)
	c.JSON(dtos.Login{Code: http.StatusOK, Message: mes})

}
