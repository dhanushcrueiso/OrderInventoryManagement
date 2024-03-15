package handlers

import (
	"OrderInventoryManagement/internal/dtos"
	"OrderInventoryManagement/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
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
	var user dtos.User
	if err := c.BodyParser(&user); err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hashedPassword)

	if err := services.SaveUser(c, user); err != nil {
		return
	}
	c.JSON(dtos.Response{Code: http.StatusOK, Message: "signup successfull"})
}
