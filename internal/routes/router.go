package routes

import (
	"OrderInventoryManagement/internal/handlers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	// Routes that require authentication middleware
	api := app.Group("/v1")
	api.Get("/ping", handlers.Ping)

}
