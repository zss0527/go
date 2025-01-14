package repository

import (
	"gin_gorm_demo/model"
	"gin_gorm_demo/pkg/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (ur *UserRepository) CreateUser(user model.User) error {
	return ur.db.Create(&user).Error
}

func (ur *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := ur.db.Where("username =?", username).First(&user).Error
	return user, err
}

// func (ur *UserRepository) CreateUser(user model.User) error {

// }
