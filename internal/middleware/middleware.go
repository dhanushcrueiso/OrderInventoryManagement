package middleware

// func ValidateToken() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		tokenValue := c.Get("Authorization")
// 		if tokenValue == "" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "Missing token",
// 			})
// 		}
// 		// Validate token against database
// 		token, err := GetTokenByValue(tokenValue)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "Invalid token",
// 			})
// 		}
// 		// Optionally, you can attach user information to the context for downstream handlers
// 		c.Locals("userID", token.UserID)
// 		return c.Next()
// 	}
// }
