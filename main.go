package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pepsi/go-fiber/app/order_api/entities"
	"github.com/pepsi/go-fiber/config"
	routes "github.com/pepsi/go-fiber/router"
)

func main() {

	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db, err := config.InitDB()
	if err != nil {
		fmt.Printf("Error initializing database: %s\n", err)
		return
	}

	db.AutoMigrate(&entities.Order{}, &entities.UploadFile{})

	routes.RegisterOrderRoutes(app, db)

	app.Listen(":8080")

}
