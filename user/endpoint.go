package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type createUserRequest struct {
	Name string
}
type createUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type deleteUserRequest struct {
	UserId string
}
type deleteUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type updateUserRequest struct {
	User User
}
type updateUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type getUserByIdRequest struct {
	UserId string
}
type getUserByIdResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type listUsersRequest struct{}
type listUsersResponse struct {
	User *[]User `json:"users,omitempty"`
	Err  error   `json:"error,omitempty"`
}

func makeHttpCreateUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createUserRequest)
		user, err := s.CreateUser(ctx, req.Name)
		return createUserResponse{&user, err}, nil
	}
}

func makeHttpDeleteUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteUserRequest)
		user, err := s.DeleteUser(ctx, req.UserId)
		return deleteUserResponse{&user, err}, nil
	}
}

func makeHttpUpdateUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateUserRequest)
		user, err := s.UpdateUser(ctx, req.User)
		return updateUserResponse{&user, err}, nil
	}
}

func makeHttpGetUserByIdEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUserByIdRequest)
		user, err := s.GetUserById(ctx, req.UserId)
		return getUserByIdResponse{&user, err}, nil
	}
}

func makeHttpListUsersEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		user, err := s.ListUsers(ctx)
		return listUsersResponse{&user, err}, nil
	}
}
