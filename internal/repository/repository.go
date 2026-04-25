package repository

import (
	"context"

	"github.com/prabowoteguh/belajar-vibe-code/internal/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user *model.User) error
}
