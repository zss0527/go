package service

import (
	"fmt"
	"gin_gorm_demo/internal/repository"
	"gin_gorm_demo/model"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: *repository.NewUserRepository(),
	}
}

func (us *UserService) Register(user model.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *UserService) Login(username, password string) error {
	user, err := us.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user.Password != password {
		return fmt.Errorf("密码错误")
	}
	return nil
}

func (us *UserService) AddUser(user model.User) error {
	err := us.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
