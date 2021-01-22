package todo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type todoCreateTodoRequest struct {
	Text,
	UserId string
}
type todoCreateTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type todoDeleteTodoRequest struct {
	TodoId string
}
type todoDeleteTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type todoUpdateTodoRequest struct {
	Todo Todo
}
type todoUpdateTodoResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type todoGetTodoByIdRequest struct {
	TodoId string
}
type todoGetTodoByIdResponse struct {
	Todo *Todo `json:"todo,omitempty"`
	Err  error `json:"error,omitempty"`
}

type todoListTodosByUserIdRequest struct {
	UserId string
}
type todoListTodosByUserIdResponse struct {
	Todo *[]Todo `json:"todo,omitempty"`
	Err  error   `json:"error,omitempty"`
}

func makeHttpTodoCreateEndpoint(s TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todoCreateTodoRequest)
		t, err := s.CreateTodo(ctx, req.Text, req.UserId)
		return todoCreateTodoResponse{&t, err}, nil
	}
}

func makeHttpTodoDeleteEndpoint(s TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todoDeleteTodoRequest)
		t, err := s.DeleteTodo(ctx, req.TodoId)
		return todoDeleteTodoResponse{&t, err}, nil
	}
}

func makeHttpTodoUpdateEndpoint(s TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todoUpdateTodoRequest)
		t, err := s.UpdateTodo(ctx, req.Todo)
		return todoUpdateTodoResponse{&t, err}, nil
	}
}

func makeHttpTodoGetTodoByIdEndpoint(s TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todoGetTodoByIdRequest)
		t, err := s.GetTodoById(ctx, req.TodoId)
		return todoGetTodoByIdResponse{&t, err}, nil
	}
}

func makeHttpTodoListTodosByUserIdEndpoint(s TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todoListTodosByUserIdRequest)
		t, err := s.ListTodosByUserId(ctx, req.UserId)
		return todoListTodosByUserIdResponse{&t, err}, nil
	}
}
