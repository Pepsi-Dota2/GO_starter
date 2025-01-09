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

	//user login
	userHandler := config.SetupUserDependencies(db)

	app.Post("/user/register", userHandler.CreateUserRegister)
	app.Post("/user/login", userHandler.UserLogin)
	app.Get("/user", userHandler.GetALlUser)
	app.Get("/user/:id", userHandler.GetUserById)
	app.Put("/user/update/:id", userHandler.UpdateUser)

}
