package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prabowoteguh/belajar-vibe-code/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) GetUsers(ctx context.Context) ([]model.UserResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.UserResponse), args.Error(1)
}

func (m *mockUserService) CreateUser(ctx context.Context, req model.CreateUserRequest) (model.UserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(model.UserResponse), args.Error(1)
}

func TestHealthCheck(t *testing.T) {
	h := NewUserHandler(nil)
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	h.HealthCheck(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
