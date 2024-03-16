package handlers

import (
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)

func Ping(c *fiber.Ctx) {
	// Parse request body and call service
	fmt.Println("wotking now ")

	err := services.CheckName(c, "275d2a3b-c827-40d4-8671-f9fb4d39e966")
	if err != nil {
		log.Fatalln("unable to fetch data ")
	}

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

	if err := services.Login(c, user); err != nil {
		return
	}
	c.JSON(dtos.Response{Code: http.StatusOK, Message: "login successfull"})
}
