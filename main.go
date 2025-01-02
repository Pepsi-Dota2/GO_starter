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
	configDatabase := config.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "myuser",
		Password: "mypassword",
		DBName:   "mydatabase",
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
