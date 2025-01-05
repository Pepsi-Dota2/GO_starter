package config

import (
	"github.com/pepsi/go-fiber/adapters"
	"github.com/pepsi/go-fiber/usecases"
	"gorm.io/gorm"
)

func SetupOrderDependencies(db *gorm.DB) *adapters.HttpOrderHandler {
	orderRepository := adapters.NewGormOrderRepositoryImpl(db)
	orderUsecase := usecases.NewOrderService(orderRepository)
	orderController := adapters.NewHttpOrderHandler(orderUsecase)
	return orderController
}
