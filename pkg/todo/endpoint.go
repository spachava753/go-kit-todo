package todo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type CreateTodoRequest struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}
type CreateTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
}

type DeleteTodoRequest struct {
	TodoId string `json:"todo_id"`
}
type DeleteTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
}

type UpdateTodoRequest struct {
	Todo Todo
}
type UpdateTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
}

type GetTodoByIdRequest struct {
	TodoId string
}
type GetTodoByIdResponse struct {
	Todo *Todo `json:"todo,omitempty"`
}

type GetTodosByUserIdRequest struct {
	UserId string `json:"user_id"`
}
type GetTodosByUserIdResponse struct {
	Todo *[]Todo `json:"todos,omitempty"`
}

func MakeCreateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateTodoRequest)
		t, err := s.CreateTodo(ctx, req.Text, req.UserId)
		return CreateTodoResponse{&t}, err
	}
}

func MakeDeleteTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteTodoRequest)
		t, err := s.DeleteTodo(ctx, req.TodoId)
		return DeleteTodoResponse{&t}, err
	}
}

func MakeUpdateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateTodoRequest)
		t, err := s.UpdateTodo(ctx, req.Todo)
		return UpdateTodoResponse{&t}, err
	}
}

func MakeGetTodoByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetTodoByIdRequest)
		t, err := s.GetTodoById(ctx, req.TodoId)
		return GetTodoByIdResponse{&t}, err
	}
}

func MakeGetTodosByUserIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetTodosByUserIdRequest)
		t, err := s.ListTodosByUserId(ctx, req.UserId)
		return GetTodosByUserIdResponse{&t}, err
	}
}
