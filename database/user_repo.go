package database

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v4"
	twitter "github.com/kaungmyathan22/golang-graphql-social-network"
)

type UserRepo struct {
	DB *DB
}

func (repo UserRepo) Create(ctx context.Context, user twitter.User) (twitter.User, error) {
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return twitter.User{}, fmt.Errorf("error starting the transaction")
	}
	err = tx.Rollback(ctx)
	if err != nil {
		return twitter.User{}, err
	}
	user, err = createUser(ctx, tx, user)
	if err != nil {
		return twitter.User{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return twitter.User{}, fmt.Errorf("error commiting")
	}
	return user, nil
}

func createUser(ctx context.Context, tx pgx.Tx, user twitter.User) (twitter.User, error) {
	query := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING *;`

	u := twitter.User{}

	if err := pgxscan.Get(ctx, tx, &u, query, user.Email, user.Username, user.Password); err != nil {
		return twitter.User{}, fmt.Errorf("error insert: %v", err)
	}

	return u, nil
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
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	user := twitter.User{}
	if err := pgxscan.Get(ctx, repo.DB.Pool, &user, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.User{}, twitter.ErrNotFound
		}
		return twitter.User{}, fmt.Errorf("error select: %v", err)
	}
	return user, nil
}

func (repo UserRepo) GetByID(ctx context.Context, id string) (twitter.User, error) {
	query := `SELECT * FROM users WHERE id = $1 LIMIT 1;`
	user := twitter.User{}
	if err := pgxscan.Get(ctx, repo.DB.Pool, &user, query, id); err != nil {
		if pgxscan.NotFound(err) {
			return twitter.User{}, twitter.ErrNotFound
		}
		return twitter.User{}, fmt.Errorf("error select: %v", err)
	}
	return user, nil
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
