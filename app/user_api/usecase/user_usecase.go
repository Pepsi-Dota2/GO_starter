package userUseCase

import (
	entities_user "github.com/pepsi/go-fiber/app/user_api/entities"
	"github.com/pepsi/go-fiber/app/user_api/repository"
)

type UserUseCase interface {
	CreateUserLogin(user *entities_user.User) error
	UpdateUser(userId uint, user *entities_user.User) error
	GetAllUser() ([]entities_user.User, error)
	GetUserById(userId uint) (*entities_user.User, error)
	DeleteUser(userId uint) error
	LoginUser(user *entities_user.User) (*entities_user.User, error)
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserUseCase {
	return &UserService{repository: repo}
}

func (s *UserService) CreateUserLogin(user *entities_user.User) error {
	return s.repository.Create(user)
}

func (s *UserService) LoginUser(user *entities_user.User) (*entities_user.User, error) {
	return s.repository.UserLogin(user)
}

func (s *UserService) UpdateUser(userId uint, user *entities_user.User) error {
	return s.repository.Update(userId, *user)
}

func (s *UserService) GetAllUser() ([]entities_user.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) GetUserById(userId uint) (*entities_user.User, error) {
	return s.repository.GetById(userId)
}

func (s *UserService) DeleteUser(userId uint) error {
	return s.repository.Delete(userId)
}
