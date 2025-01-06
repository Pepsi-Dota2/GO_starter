package adapters

import (
	"github.com/pepsi/go-fiber/app/order_api/entities"
	"github.com/pepsi/go-fiber/app/order_api/repository"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepositoryImpl(db *gorm.DB) repository.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order *entities.Order) error {
	return r.db.Create(&order).Error
}

func (r *GormOrderRepository) Update(id uint, updatedOrder entities.Order) error {
	var order entities.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return err
	}

	return r.db.Updates(updatedOrder).Error
}

func (r *GormOrderRepository) GetAll() ([]entities.Order, error) {
	var orders []entities.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *GormOrderRepository) UploadFile(file entities.UploadFile) error {
	return r.db.Create(&file).Error
}
