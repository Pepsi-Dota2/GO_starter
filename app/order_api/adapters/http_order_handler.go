package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/app/order_api/entities"
	"github.com/pepsi/go-fiber/app/order_api/usecases"
)

type HttpOrderHandler struct {
	orderUsecase usecases.OrderUsecase
}

func NewHttpOrderHandler(useCase usecases.OrderUsecase) *HttpOrderHandler {
	return &HttpOrderHandler{
		orderUsecase: useCase,
	}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.orderUsecase.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *HttpOrderHandler) UpdateOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Parse the request body to get the updated order data
	var updatedOrder entities.Order
	if err := c.BodyParser(&updatedOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call the service or repository to update the order
	updatedOrder.ID = uint(orderId)
	err = h.orderUsecase.UpdateOrder(uint(orderId), updatedOrder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedOrder)

}

func (h *HttpOrderHandler) GetAllOrder(c *fiber.Ctx) error {
	orders, err := h.orderUsecase.GetAllOrder()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(orders)
}

func (h *HttpOrderHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = h.orderUsecase.UploadFile(entities.UploadFile{
		File: file.Filename,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"file":    file.Filename,
	})
}
