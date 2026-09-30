package service

import (
	"context"
	"errors"
	"rest_api/internal/model"
	"rest_api/internal/repository"
)

type UserService struct{
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository)*UserService{
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(
	ctx context.Context,
	email string,
	password string,
) (model.User, error) {

	if email == "" {
		return model.User{}, errors.New("email is required")
	}

	if password == "" {
		return model.User{}, errors.New("password is required")
	}

	// Здесь позже захешируем password

	user := model.User{
		Email:        email,
		PasswordHash: password, // ВРЕМЕННО! Потом здесь будет hash
	}

	return s.repo.Create(ctx, user)
}
