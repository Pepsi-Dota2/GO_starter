package adapters

import (
	"log"

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
	log.Printf("Saving order: %+v\n", order)
	if err := r.db.Create(&order).Error; err != nil {
		log.Printf("Failed to save order: %v\n", err)
		return err
	}
	log.Printf("Order saved successfully: %+v\n", order)
	return nil
}

func (r *GormOrderRepository) Update(id uint, updatedOrder entities.Order) error {
	log.Printf("Updating order with ID: %d\n", id)
	var order entities.Order
	if err := r.db.First(&order, id).Error; err != nil {
		log.Printf("Order with ID %d not found: %v\n", id, err)
		return err
	}

	if err := r.db.Updates(updatedOrder).Error; err != nil {
		log.Printf("Failed to update order with ID %d: %v\n", id, err)
		return err
	}
	log.Printf("Order with ID %d updated successfully\n", id)
	return nil
}

func (r *GormOrderRepository) GetAll() ([]entities.Order, error) {
	log.Println("Fetching all orders")
	var orders []entities.Order
	if err := r.db.Find(&orders).Error; err != nil {
		log.Printf("Failed to fetch orders: %v\n", err)
		return nil, err
	}
	log.Printf("Fetched %d orders\n", len(orders))
	return orders, nil
}

func (r *GormOrderRepository) UploadFile(file entities.UploadFile) error {
	log.Printf("Uploading file: %+v\n", file)
	if err := r.db.Create(&file).Error; err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		return err
	}
	log.Printf("File uploaded successfully: %+v\n", file)
	return nil
}

func (r *GormOrderRepository) GetById(id uint) (*entities.Order, error) {
	log.Printf("Fetching order with ID: %d\n", id)
	var order entities.Order
	if err := r.db.First(&order, id).Error; err != nil {
		log.Printf("Order with ID %d not found: %v\n", id, err)
		return nil, err
	}
	log.Printf("Order with ID %d fetched successfully: %+v\n", id, order)
	return &order, nil
}
