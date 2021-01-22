package todo

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type basicLoggingService struct {
	logger log.Logger
	Service
}

func (s *basicLoggingService) CreateTodo(ctx context.Context, text string, userId string) (t Todo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "create", "text", text, "userId", userId, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.TodoService.CreateTodo(ctx, text, userId)
}

func (s *basicLoggingService) DeleteTodo(ctx context.Context, todoId string) (t Todo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "delete", "todoId", todoId, "text", t.Text, "done", t.Done, "userId", t.UserId, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.TodoService.DeleteTodo(ctx, todoId)
}

func (s *basicLoggingService) UpdateTodo(ctx context.Context, todo Todo) (t Todo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "update", "updatedText", t.Text, "updatedDone", t.Done, "updatedUserId", t.UserId, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.TodoService.UpdateTodo(ctx, todo)
}

func (s *basicLoggingService) GetTodoById(ctx context.Context, todoId string) (t Todo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "get", "todoId", todoId, "text", t.Text, "done", t.Done, "userId", t.UserId, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.TodoService.GetTodoById(ctx, todoId)
}

func (s *basicLoggingService) ListTodosByUserId(ctx context.Context, userId string) (ts []Todo, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "list", "len", len(ts), "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.TodoService.ListTodosByUserId(ctx, userId)
}

func NewBasicLoggingService(logger log.Logger, s Service) Service {
	return &basicLoggingService{
		logger:      logger,
		TodoService: s,
	}
}
