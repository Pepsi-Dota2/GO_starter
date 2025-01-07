package usecases

import (
	"github.com/pepsi/go-fiber/app/order_api/entities"
	"github.com/pepsi/go-fiber/app/order_api/repository"
)

type OrderUsecase interface {
	CreateOrder(order *entities.Order) error
	UpdateOrder(orderId uint, order entities.Order) error
	GetAllOrder() ([]entities.Order, error)
	UploadFile(file entities.UploadFile) error
	GetOrderById(orderId uint) (*entities.Order, error)
}
type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderUsecase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *entities.Order) error {
	return s.repo.Save(order)
}

func (s *OrderService) UpdateOrder(orderId uint, order entities.Order) error {
	return s.repo.Update(uint(orderId), order)
}

func (s *OrderService) GetAllOrder() ([]entities.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) UploadFile(file entities.UploadFile) error {
	return s.repo.UploadFile(file)
}

func (s *OrderService) GetOrderById(orderId uint) (*entities.Order, error) {
	return s.repo.GetById(orderId)
}
