package repository

import (
	"errors"
	models "go_rest_api/model"
)

type UserRepo interface {
	CreateUser(user *models.User) bool
	FindUser(id string) (*models.User, error)
}

var (
	listUser = make([]*models.User, 0)
)

func CreateUser(user *models.User) bool {
	if user.Id != "" && user.Email != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return false
}

func FindUser(id string) (*models.User, error) {
	for _, user := range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("User không tồn tại")
}
