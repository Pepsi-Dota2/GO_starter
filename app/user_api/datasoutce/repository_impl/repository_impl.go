package repositoryimpl

import (
	"log"

	entities_user "github.com/pepsi/go-fiber/app/user_api/entities"
	"github.com/pepsi/go-fiber/app/user_api/repository"
	"gorm.io/gorm"
)

type GormUserRepositoryImpl struct {
	db *gorm.DB
}

func NewGormUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &GormUserRepositoryImpl{db: db}
}

func (r *GormUserRepositoryImpl) Create(user *entities_user.User) error {
	log.Printf("Saving user: %+v\n", user)
	if err := r.db.Create(&user).Error; err != nil {
		log.Printf("Failed to save user: %v\n", err)
		return err
	}
	log.Printf("User saved successfully: %+v\n", user)
	return nil
}

func (r *GormUserRepositoryImpl) GetAll() ([]entities_user.User, error) {
	var users []entities_user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepositoryImpl) GetById(id uint) (*entities_user.User, error) {
	var user entities_user.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepositoryImpl) Update(id uint, updatedUser entities_user.User) error {
	if err := r.db.Where("id = ?", id).Updates(&updatedUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormUserRepositoryImpl) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&entities_user.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormUserRepositoryImpl) UserLogin(user *entities_user.User) (*entities_user.User, error) {
	var foundUser entities_user.User

	// Use AND to check both username AND password match
	result := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&foundUser)
	if result.Error != nil {
		log.Printf("Login error: %v", result.Error)
		log.Printf("Attempted username: %s", user.Username)
		return nil, result.Error
	}

	return &foundUser, nil
}
