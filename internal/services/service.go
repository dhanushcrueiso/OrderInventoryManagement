package services

import (
	"OrderInventoryManagement/internal/database/daos"
	"fmt"

	"github.com/gofiber/fiber"
)

func CheckName(c *fiber.Ctx, req string) error {

	res, err := daos.GetName(c, req)
	if err != nil {
		return err
	}
	fmt.Println("res", res)
	return nil
}
