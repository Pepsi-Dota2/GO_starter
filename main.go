package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/pepsi/go-fiber/config"
	"github.com/pepsi/go-fiber/entities"
	routes "github.com/pepsi/go-fiber/router"
)

func main() {

	app := fiber.New()
	db, err := config.InitDB()
	if err != nil {
		fmt.Printf("Error initializing database: %s\n", err)
		return
	}

	db.AutoMigrate(&entities.Order{})

	routes.RegisterOrderRoutes(app, db)

	app.Listen(":8080")

}
