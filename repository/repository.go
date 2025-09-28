package repository

import (
	"suretigai/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id int64) (*entity.User, error)
}

