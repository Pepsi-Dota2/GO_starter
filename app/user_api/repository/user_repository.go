package repository

import entities_user "github.com/pepsi/go-fiber/app/user_api/entities"

type UserRepository interface {
	Create(user *entities_user.User) error
	Update(id uint, updatedUser entities_user.User) error
	GetAll() ([]entities_user.User, error)
	GetById(id uint) (*entities_user.User, error)
	Delete(id uint) error
	UserLogin(user *entities_user.User) (*entities_user.User, error)
}
