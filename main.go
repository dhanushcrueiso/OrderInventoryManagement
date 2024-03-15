package main

import (
	"OrderInventoryManagement/config"
	"OrderInventoryManagement/internal/database/db"
	"log"
	"os"
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
	// r := routes.GetRouter()
	// constants.Logger.Info("Listening to Port: " + cnf.Port)
	// r.Run(":" + cnf.Port)
}
