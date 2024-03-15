package main

import (
	"OrderInventoryManagement/config"
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {

	env := "dev"

	var file *os.File
	var err error

	file, err = os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}

	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)

	db.Init(&db.Config{
		URL:       cnf.DatabaseURL,
		MaxDBConn: cnf.MaxDBConn,
	})

	//For Initialising Exchange Rates

	// ctx := context.Background()
	app := fiber.New()

	routes.SetupRoutes(app)
	fmt.Printf("Server is running on port %s\n", cnf.Port)

	// r := routes.GetRouter()
	// constants.Logger.Info("Listening to Port: " + cnf.Port)
	// r.Run(":" + cnf.Port)
	app.Listen(cnf.Port)
}
