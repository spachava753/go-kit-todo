package services

import "context"

type User struct {
	Id,
	Name string
}

type UserService interface {
	CreateUser(ctx context.Context, name string) (User, error)
	DeleteUser(ctx context.Context, userId string) (User, error)
	GetUserById(ctx context.Context, userId string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}
