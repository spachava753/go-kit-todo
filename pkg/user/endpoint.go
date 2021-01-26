package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type CreateUserRequest struct {
	Name string
}
type CreateUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	UserId string
}
type DeleteUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type UpdateUserRequest struct {
	User User
}
type UpdateUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type GetUserByIdRequest struct {
	UserId string
}
type GetUserByIdResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type ListUsersRequest struct{}
type ListUsersResponse struct {
	Users []User `json:"users,omitempty"`
	Err   error  `json:"error,omitempty"`
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		user, err := s.CreateUser(ctx, req.Name)
		return CreateUserResponse{&user, err}, nil
	}
}

func MakeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteUserRequest)
		user, err := s.DeleteUser(ctx, req.UserId)
		return DeleteUserResponse{&user, err}, nil
	}
}

func MakeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateUserRequest)
		user, err := s.UpdateUser(ctx, req.User)
		return UpdateUserResponse{&user, err}, nil
	}
}

func MakeGetUserByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserByIdRequest)
		user, err := s.GetUserById(ctx, req.UserId)
		return GetUserByIdResponse{&user, err}, nil
	}
}

func MakeListUsersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, err := s.ListUsers(ctx)
		return ListUsersResponse{user, err}, nil
	}
}
