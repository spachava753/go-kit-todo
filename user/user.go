package user

import (
	"context"
	"fmt"
	"github.com/oklog/ulid/v2"
	"sync"
	"time"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	CreateUser(ctx context.Context, name string) (User, error)
	DeleteUser(ctx context.Context, userId string) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	GetUserById(ctx context.Context, userId string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}

type inMemUserService struct {
	entropy *ulid.MonotonicEntropy
	time    time.Time
	users   sync.Map
}

func (s *inMemUserService) UpdateUser(ctx context.Context, user User) (User, error) {
	if _, ok := s.users.Load(user.Id); !ok {
		return User{}, fmt.Errorf("user with id %s does not exist", user.Id)
	}
	s.users.Store(user.Id, user)
	return user, nil
}

func (s *inMemUserService) CreateUser(ctx context.Context, name string) (User, error) {
	id := ulid.MustNew(ulid.Timestamp(s.time), s.entropy).String()
	user := User{id, name}
	s.users.Store(id, user)
	return user, nil
}

func (s *inMemUserService) DeleteUser(ctx context.Context, userId string) (User, error) {
	if v, ok := s.users.LoadAndDelete(userId); ok {
		return v.(User), nil
	}
	return User{}, fmt.Errorf("could not delete user with id %s", userId)
}

func (s *inMemUserService) GetUserById(ctx context.Context, userId string) (User, error) {
	if v, ok := s.users.Load(userId); ok {
		return v.(User), nil
	}
	return User{}, fmt.Errorf("could not find user with id %s", userId)
}

func (s *inMemUserService) ListUsers(ctx context.Context) ([]User, error) {
	var userList []User
	s.users.Range(func(_, value interface{}) bool {
		userList = append(userList, value.(User))
		return true
	})
	return userList, nil
}

func NewMemUserService(entropy *ulid.MonotonicEntropy, t time.Time) Service {
	return &inMemUserService{entropy, t, sync.Map{}}
}
