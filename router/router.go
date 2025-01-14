package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/config"
	"github.com/pepsi/go-fiber/middleware"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(app *fiber.App, db *gorm.DB) {
	orderHandler := config.SetupOrderDependencies(db)
	userHandler := config.SetupUserDependencies(db)

	// Public routes (no authentication required)
	app.Post("/user/login", userHandler.UserLogin)
	app.Post("/user/register", userHandler.CreateUserRegister)
	app.Post("/user/logout", userHandler.LogoutUser)

	// Protected routes (authentication required)
	api := app.Group("/")
	api.Use(middleware.AuthRequired) // Add this line to apply the middleware

	// Protected order routes
	api.Post("/orders", orderHandler.CreateOrder)
	api.Put("/orders/:id", orderHandler.UpdateOrder)
	api.Get("/orders", orderHandler.GetAllOrder)
	api.Post("/orders/upload", orderHandler.UploadFile)
	api.Get("/orders/:id", orderHandler.GetOrderById)

	// Protected user routes
	api.Get("/user", userHandler.GetALlUser)
	api.Get("/user/:id", userHandler.GetUserById)
	api.Put("/user/update/:id", userHandler.UpdateUser)
	api.Delete("/user/delete/:id", userHandler.DeleteUser)

}
