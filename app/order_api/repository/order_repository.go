package repository

import "github.com/pepsi/go-fiber/app/order_api/entities"

type OrderRepository interface {
	Save(order *entities.Order) error
	Update(id uint, updatedOrder entities.Order) error
	GetAll() ([]entities.Order, error)
	GetById(id uint) (*entities.Order, error)
	UploadFile(file entities.UploadFile) error
}
