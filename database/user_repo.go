package database

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	twitter "github.com/kaungmyathan22/golang-graphql-social-network"
)

type UserRepo struct {
	DB *DB
}

func (repo UserRepo) Create(ctx context.Context, user twitter.User) (twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo UserRepo) GetByUsername(ctx context.Context, username string) (twitter.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`
	user := twitter.User{}
	if err := pgxscan.Get(ctx, repo.DB.Pool, &user, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.User{}, twitter.ErrNotFound
		}
		return twitter.User{}, fmt.Errorf("error select: %v", err)
	}
	return user, nil
}

func (repo UserRepo) GetByEmail(ctx context.Context, email string) (twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo UserRepo) GetByID(ctx context.Context, id string) (twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo UserRepo) GetByIds(ctx context.Context, ids []string) ([]twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
