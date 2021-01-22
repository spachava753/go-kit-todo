package services

import "context"

type Todo struct {
	Id,
	Text string
	Done   bool
	UserId string
}

type TodoService interface {
	CreateTodo(ctx context.Context, text string, userId string) (Todo, error)
	DeleteTodo(ctx context.Context, todoId string) (Todo, error)
	UpdateTodo(ctx context.Context, todoId, text string, done bool) (Todo, error)
	GetTodoById(ctx context.Context, todoId string) (Todo, error)
	ListTodosByUserId(ctx context.Context, userId string) ([]Todo, error)
}
