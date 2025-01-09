package config

import (
	"github.com/pepsi/go-fiber/app/order_api/adapters"
	"github.com/pepsi/go-fiber/app/order_api/usecases"
	"github.com/pepsi/go-fiber/app/user_api/datasoutce/data"
	repositoryimpl "github.com/pepsi/go-fiber/app/user_api/datasoutce/repository_impl"
	userUseCase "github.com/pepsi/go-fiber/app/user_api/usecase"
	"gorm.io/gorm"
)

func SetupOrderDependencies(db *gorm.DB) *adapters.HttpOrderHandler {
	orderRepository := adapters.NewGormOrderRepositoryImpl(db)
	orderUsecase := usecases.NewOrderService(orderRepository)
	orderController := adapters.NewHttpOrderHandler(orderUsecase)
	return orderController
}

func SetupUserDependencies(db *gorm.DB) *data.HttpUserHandler {
	userRepository := repositoryimpl.NewGormUserRepositoryImpl(db)
	userUsecase := userUseCase.NewUserService(userRepository)
	userController := data.NewUserHttpHandler(userUsecase)
	return userController
}
