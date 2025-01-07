package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/config"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(app *fiber.App, db *gorm.DB) {
	orderHandler := config.SetupOrderDependencies(db)

	app.Post("/orders", orderHandler.CreateOrder)
	app.Put("/orders/:id", orderHandler.UpdateOrder)
	app.Get("/orders", orderHandler.GetAllOrder)
	app.Post("/orders/upload", orderHandler.UploadFile)
	app.Get("/orders/:id", orderHandler.GetOrderById)
}
