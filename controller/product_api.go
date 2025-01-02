package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/model"
	"github.com/pepsi/go-fiber/service"
)

func GetProductByIdHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := service.GetProductById(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func CreateProductHandler(c *fiber.Ctx) error {
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := service.CreateProduct(product)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func UpdateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Invalid product ID: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Printf("Invalid request body: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := service.UpdateProduct(productId, p)

	if err != nil {
		log.Printf("Failed to update product: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func DeleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("Invalid product ID: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = service.DeleteProduct(productId)
	if err != nil {
		log.Printf("Failed to delete product: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func GetAllProductsHandler(c *fiber.Ctx) error {
	products, err := service.GetAllProducts()
	if err != nil {
		log.Printf("Failed to get products: %v\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(products)
}
