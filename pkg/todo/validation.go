package todo

import (
	"context"
	"github.com/spachava/go-kit-todo/pkg/user"
)

type userValidationService struct {
	userSvc user.Service
	Service
}

func (s *userValidationService) CreateTodo(ctx context.Context, text string, userId string) (t Todo, err error) {
	if _, err := s.userSvc.GetUserById(ctx, userId); err != nil {
		return Todo{}, err
	}
	return s.Service.CreateTodo(ctx, text, userId)
}

func (s *userValidationService) DeleteTodo(ctx context.Context, todoId string) (t Todo, err error) {
	return s.Service.DeleteTodo(ctx, todoId)
}

func (s *userValidationService) UpdateTodo(ctx context.Context, todo Todo) (t Todo, err error) {
	if _, err := s.userSvc.GetUserById(ctx, todo.UserId); err != nil {
		return Todo{}, err
	}
	return s.Service.UpdateTodo(ctx, todo)
}

func (s *userValidationService) GetTodoById(ctx context.Context, todoId string) (t Todo, err error) {
	return s.Service.GetTodoById(ctx, todoId)
}

func (s *userValidationService) ListTodosByUserId(ctx context.Context, userId string) (ts []Todo, err error) {
	if _, err := s.userSvc.GetUserById(ctx, userId); err != nil {
		return nil, err
	}
	return s.Service.ListTodosByUserId(ctx, userId)
}

func NewUserValidationService(todoSvc Service, userSvc user.Service) Service {
	return &userValidationService{
		userSvc,
		todoSvc,
	}
}
