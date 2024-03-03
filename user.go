package twitter

import (
	"context"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, user User) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
	GetByIds(ctx context.Context, ids []string) ([]User, error)
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}