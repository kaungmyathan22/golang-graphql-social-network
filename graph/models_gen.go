// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"time"
)

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	User        *User  `json:"user"`
}

type CreateTweetInput struct {
	Body string `json:"body"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterInput struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type Tweet struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	User      *User     `json:"user"`
	UserID    string    `json:"userId"`
	ParentID  *string   `json:"parentId,omitempty"`
	Replies   []*Tweet  `json:"replies"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
