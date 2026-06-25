package service

import (
	"context"
	"errors"
	"strings"

	"ci-cd-go-learn/internal/model"
	"ci-cd-go-learn/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, req model.CreateUserRequest) (model.User, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if req.Name == "" {
		return model.User{}, errors.New("name is required")
	}
	if req.Email == "" {
		return model.User{}, errors.New("email is required")
	}

	return s.repo.Create(ctx, req)
}