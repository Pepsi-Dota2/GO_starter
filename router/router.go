package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/product/:id", controller.GetProductByIdHandler)
	app.Post("/product/create", controller.CreateProductHandler)
	app.Put("/product/update/:id", controller.UpdateProductHandler)
	app.Delete("/product/delete/:id", controller.DeleteProductHandler)
	app.Get("/products", controller.GetAllProductsHandler)
}
