package usecases

import (
	"github.com/pepsi/go-fiber/entities"
	"github.com/pepsi/go-fiber/repository"
)

type OrderUsecase interface {
	CreateOrder(order entities.Order) error
}
type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderUsecase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {
	return s.repo.Save(order)
}
