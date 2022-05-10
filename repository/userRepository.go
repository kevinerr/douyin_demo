package repository

import (
	"github.com/RaymondCode/simple-demo/model"
)

type UserRepository struct {
}

func (c UserRepository) CreateUser(user *model.User) error {
	err := model.DB.Create(user).Error
	return err
}

func (c UserRepository) IsExistUser(username string) (*model.User, bool) {
	var user model.User
	var count int
	model.DB.Where("user_name=?", username).First(&user).Count(&count)
	if count == 1 {
		return &user, true
	}
	return &user, false
}

func (c UserRepository) SelectById(userId string) (*model.User, error) {
	var user model.User
	if err := model.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
