package config

import (
	"github.com/pepsi/go-fiber/app/order_api/adapters"
	"github.com/pepsi/go-fiber/app/order_api/usecases"
	"gorm.io/gorm"
)

func SetupOrderDependencies(db *gorm.DB) *adapters.HttpOrderHandler {
	orderRepository := adapters.NewGormOrderRepositoryImpl(db)
	orderUsecase := usecases.NewOrderService(orderRepository)
	orderController := adapters.NewHttpOrderHandler(orderUsecase)
	return orderController
}
