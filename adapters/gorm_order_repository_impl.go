package adapters

import (
	"github.com/pepsi/go-fiber/entities"
	"github.com/pepsi/go-fiber/repository"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepositoryImpl(db *gorm.DB) repository.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order entities.Order) error {
	return r.db.Create(&order).Error
}
