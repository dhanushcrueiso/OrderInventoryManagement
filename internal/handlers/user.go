package handlers

import (
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

func Ping(c *fiber.Ctx) {
	// Parse request body and call service
	fmt.Println("wotking now ")

	// err := services.CheckName(c, "275d2a3b-c827-40d4-8671-f9fb4d39e966")
	// if err != nil {
	// 	log.Fatalln("unable to fetch data ")
	// }
	c.JSON("Pong")
}

func Signup(c *fiber.Ctx) {
	fmt.Println("here:")
	var user dtos.User
	if err := c.BodyParser(&user); err != nil {
		return
	}
	fmt.Println("checking the parser", user)

	if err := services.SaveUser(c, user); err != nil {
		return
	}
	c.JSON(dtos.Response{Code: http.StatusOK, Message: "signup successfull"})
}

func Login(c *fiber.Ctx) {
	var user dtos.User
	if err := c.BodyParser(&user); err != nil {
		return
	}
	fmt.Println("checking the parser", user)

	res, err := services.Login(c, user)
	if err != nil {
		return
	}
	c.JSON(dtos.Login{AccountInfo: *res, Code: http.StatusOK, Message: "login successfull"})
}

func MostSoldProducts(c *fiber.Ctx) {

	res, err := services.MostSoldProducts(c)
	if err != nil {
		return
	}
	c.JSON(res)
}
