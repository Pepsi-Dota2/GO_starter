package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/pepsi/go-fiber/config"
	"github.com/pepsi/go-fiber/router"
	"github.com/pepsi/go-fiber/service"
)

func main() {

	configDatabase, err := config.LoadDBConfig()

	if err != nil {
		log.Fatalf("Error loading database configuration: %v\n", err)
	}
	db, err := config.NewDBConnection(
		configDatabase,
	)

	if err != nil {
		log.Fatalf("Error initializing database connection: %v\n", err)
	}

	service.DB = db

	defer service.DB.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":8080")

}
