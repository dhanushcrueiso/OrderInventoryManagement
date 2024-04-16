package middleware

import (
	"OrderInventoryManagement/internal/dtos"
	"fmt"
	"strings"

	"github.com/gofiber/fiber"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) {
		authHeader := c.Get("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "invalid token"})
		}

		// Split the Authorization header value to extract the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "invalid token"})
		}

		// Extract the token from the Authorization header
		token := parts[1]
		fmt.Println(token)

		// res, err := daos.GetAccountByToken(token)
		// if err != nil {
		// 	c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "no account mapped to the token"})
		// }
		// if res.ExpiresAt.After(time.Now()) {
		// 	c.JSON(dtos.Response{Code: fiber.ErrUnauthorized.Code, Message: "token expired"})
		// }
		// Proceed to the next middleware/handler
		c.Next() // No error indicates successful execution
	}
}

// tokenValue := c.Get("Authorization")
// if tokenValue == "" {
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": "Missing token",
// 	})
// }
// Validate token against database
// token, err := GetTokenByValue(tokenValue)
// if err != nil {
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": "Invalid token",
// 	})
// }
// Optionally, you can attach user information to the context for downstream handlers
// c.Locals("userID", token.UserID)
// return c.Next()
