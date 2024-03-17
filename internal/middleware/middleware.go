package middleware

import (
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/dtos"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) {
		authHeader := c.Get("Authorization")
		fmt.Println("asda", authHeader)
		// Check if the Authorization header is present
		if authHeader == "" {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "invalid token"})
			return
		}

		// Split the Authorization header value to extract the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "invalid token"})
			return
		}

		// Extract the token from the Authorization header
		token := parts[1]
		fmt.Println(token)

		res, err := daos.GetAccountByToken(token)
		fmt.Println("asad", res)
		fmt.Println("asad", res.ExpiresAt.Before(time.Now()))
		if err != nil {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "no account mapped to the token"})
			return
		}
		if res.ExpiresAt.Before(time.Now()) {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "token expired"})
			return
		}
		// Proceed to the next middleware/handler
		c.Next() // No error indicates successful execution
	}
}
