package routes

import (
	"OrderInventoryManagement/internal/handlers"
	"OrderInventoryManagement/internal/middleware"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	// Routes that require authentication middleware
	validateToken := middleware.ValidateToken()
	api := app.Group("/v1")
	api.Get("/ping", validateToken, handlers.Ping)

}
