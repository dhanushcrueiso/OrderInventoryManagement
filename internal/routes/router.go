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
	app.Get("/ping", handlers.Ping)
	api.Post("/admins/signup", handlers.Signup)
	api.Post("/admins/login", handlers.Login)
	api.Post("/customer/signup", handlers.CustomerSignup)
	api.Post("/customer/login", handlers.CustomerLogin)
	api.Get("/products", validateToken, handlers.GetProducts)
	api.Put("/admin/:aid/product", validateToken, handlers.AddProducts)
	api.Post("/customer/:cid/placeorder", validateToken, handlers.PlaceOrder)
	api.Get("/customer/:cid/getallorder", validateToken, handlers.GetAllOrders)
	api.Get("/admins/produt-analytics", validateToken, handlers.MostSoldProducts)

}
