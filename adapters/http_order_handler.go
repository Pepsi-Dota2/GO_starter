package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pepsi/go-fiber/entities"
	"github.com/pepsi/go-fiber/usecases"
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
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.orderUsecase.CreateOrder(order); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
