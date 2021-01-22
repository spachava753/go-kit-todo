package todo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type createTodoRequest struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}
type createTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type deleteTodoRequest struct {
	TodoId string `json:"todo_id"`
}
type deleteTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type updateTodoRequest struct {
	Todo Todo
}
type updateTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type getTodoByIdRequest struct {
	TodoId string
}
type getTodoByIdResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type getTodosByUserIdRequest struct {
	UserId string `json:"user_id"`
}
type getTodosByUserIdResponse struct {
	Todo *[]Todo `json:"todos,omitempty"`
	Err  error   `json:"error,omitempty"`
}

func makeHttpCreateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createTodoRequest)
		t, err := s.CreateTodo(ctx, req.Text, req.UserId)
		return createTodoResponse{&t, err}, nil
	}
}

func makeHttpDeleteTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteTodoRequest)
		t, err := s.DeleteTodo(ctx, req.TodoId)
		return deleteTodoResponse{&t, err}, nil
	}
}

func makeHttpUpdateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateTodoRequest)
		t, err := s.UpdateTodo(ctx, req.Todo)
		return updateTodoResponse{&t, err}, nil
	}
}

func makeHttpGetTodoByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getTodoByIdRequest)
		t, err := s.GetTodoById(ctx, req.TodoId)
		return getTodoByIdResponse{&t, err}, nil
	}
}

func makeHttpListTodosByUserIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getTodosByUserIdRequest)
		t, err := s.ListTodosByUserId(ctx, req.UserId)
		return getTodosByUserIdResponse{&t, err}, nil
	}
}
