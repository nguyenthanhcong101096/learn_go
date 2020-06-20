package repository

import (
	models "go_postgres/model"
)

type UserRepo interface {
	Select() ([]models.User, error)
}
