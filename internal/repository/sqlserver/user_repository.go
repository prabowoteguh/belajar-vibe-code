package sqlserver

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/prabowoteguh/belajar-vibe-code/internal/model"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	query := "SELECT id, username, email, created_at FROM users"
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (username, email, created_at) OUTPUT INSERTED.id VALUES (:username, :email, :created_at)"
	rows, err := r.db.NamedQueryContext(ctx, query, user)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
