package service

import (
	"context"
	"time"

	"github.com/prabowoteguh/belajar-vibe-code/internal/model"
	"github.com/prabowoteguh/belajar-vibe-code/internal/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]model.UserResponse, error)
	CreateUser(ctx context.Context, req model.CreateUserRequest) (model.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers(ctx context.Context) ([]model.UserResponse, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []model.UserResponse
	for _, u := range users {
		res = append(res, model.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return res, nil
}

func (s *userService) CreateUser(ctx context.Context, req model.CreateUserRequest) (model.UserResponse, error) {
	user := &model.User{
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
