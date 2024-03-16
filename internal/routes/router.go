package routes

import (
	"OrderInventoryManagement/internal/handlers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	// Routes that require authentication middleware
	api := app.Group("/v1")
	api.Get("/ping", handlers.Ping)
	api.Post("/admins/signup", handlers.Signup)
	api.Post("/admins/login", handlers.Login)
	api.Post("/customer/signup", handlers.CustomerSignup)
	api.Post("/customer/login", handlers.CustomerLogin)

}
