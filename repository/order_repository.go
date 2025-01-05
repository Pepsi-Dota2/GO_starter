package repository

import "github.com/pepsi/go-fiber/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
