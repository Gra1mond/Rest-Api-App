package service

import (
	"context"
	"errors"
	"rest_api/internal/auth"
	"rest_api/internal/model"
	"rest_api/internal/repository"

	"golang.org/x/crypto/bcrypt"
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

	hash,err:=bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err!=nil{
		return model.User{},err
	}

	user := model.User{
		Email:        email,
		PasswordHash: string(hash), 
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}