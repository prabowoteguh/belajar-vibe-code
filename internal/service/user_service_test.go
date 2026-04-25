package service

import (
	"context"
	"errors"
	"testing"

	"github.com/prabowoteguh/belajar-vibe-code/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	if user != nil {
		user.ID = 1 // simulate ID assignment
	}
	return args.Error(0)
}

func TestGetUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything).Return([]model.User{
			{ID: 1, Username: "testuser", Email: "test@example.com"},
		}, nil).Once()

		users, err := svc.GetUsers(context.Background())

		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "testuser", users[0].Username)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything).Return(nil, errors.New("db error")).Once()

		users, err := svc.GetUsers(context.Background())

		assert.Error(t, err)
		assert.Nil(t, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		req := model.CreateUserRequest{
			Username: "newuser",
			Email:    "new@example.com",
		}

		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil).Once()

		user, err := svc.CreateUser(context.Background(), req)

		assert.NoError(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "newuser", user.Username)
		mockRepo.AssertExpectations(t)
	})
}
