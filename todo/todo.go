package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"sync"
	"time"
)

type Todo struct {
	Id,
	Text string
	Done   bool
	UserId string
}

type Service interface {
	CreateTodo(ctx context.Context, text string, userId string) (Todo, error)
	DeleteTodo(ctx context.Context, todoId string) (Todo, error)
	UpdateTodo(ctx context.Context, todo Todo) (Todo, error)
	GetTodoById(ctx context.Context, todoId string) (Todo, error)
	ListTodosByUserId(ctx context.Context, userId string) ([]Todo, error)
}

type inMemTodoService struct {
	entropy *ulid.MonotonicEntropy
	time    time.Time
	todos   sync.Map
}

func (s *inMemTodoService) UpdateTodo(ctx context.Context, todo Todo) (Todo, error) {
	if _, ok := s.todos.Load(todo.Id); !ok {
		return Todo{}, fmt.Errorf("user with id %s does not exist", todo.Id)
	}
	s.todos.Store(todo.Id, todo)
	return todo, nil
}

func (s *inMemTodoService) CreateTodo(ctx context.Context, text string, userId string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("todo text cannot be empty")
	}
	id := ulid.MustNew(ulid.Timestamp(s.time), s.entropy).String()
	todo := Todo{id, text, false, userId}
	s.todos.Store(id, todo)
	return todo, nil
}

func (s *inMemTodoService) DeleteTodo(ctx context.Context, todoId string) (Todo, error) {
	if v, ok := s.todos.LoadAndDelete(todoId); ok {
		return v.(Todo), nil
	}
	return Todo{}, fmt.Errorf("could not todo delete with id %s", todoId)
}

func (s *inMemTodoService) GetTodoById(ctx context.Context, todoId string) (Todo, error) {
	if v, ok := s.todos.Load(todoId); ok {
		return v.(Todo), nil
	}
	return Todo{}, fmt.Errorf("could not find todo with id %s", todoId)
}

func (s *inMemTodoService) ListTodosByUserId(ctx context.Context, userId string) ([]Todo, error) {
	var todoList []Todo
	s.todos.Range(func(_, value interface{}) bool {
		todo := value.(Todo)
		if todo.UserId == userId {
			todoList = append(todoList, todo)
		}
		return true
	})
	return todoList, nil
}

func NewMemTodoService(entropy *ulid.MonotonicEntropy, t time.Time) Service {
	return &inMemTodoService{entropy, t, sync.Map{}}
}
