package usecase

import (
	"errors"
	"suretigai/entity"
	"suretigai/repository"
)

type UserUsecaseInterface interface {
	RegisterUser(user *entity.User) (*entity.User, error)
}

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: r,
	}
}

// ユーザー登録
func (u *UserUsecase) HandleRegisterUser(user *entity.User) (*entity.User, error) {
	// バリデーションを追
	if user.Name == "" {
		return nil, errors.New("名前は必須です")
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
