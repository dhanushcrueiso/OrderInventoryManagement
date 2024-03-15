package handlers

import (
	"OrderInventoryManagement/internal/services"
	"fmt"
	"log"

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
