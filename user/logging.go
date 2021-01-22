package user

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type basicLoggingService struct {
	logger log.Logger
	Service
}

func (s *basicLoggingService) CreateUser(ctx context.Context, name string) (u User, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "create", "name", name, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.UserService.CreateUser(ctx, name)
}

func (s *basicLoggingService) DeleteUser(ctx context.Context, userId string) (u User, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "delete", "name", u.Name, "userId", u.Id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.UserService.DeleteUser(ctx, userId)
}

func (s *basicLoggingService) UpdateUser(ctx context.Context, user User) (u User, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "update", "updatedName", user.Name, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.UserService.UpdateUser(ctx, user)
}

func (s *basicLoggingService) GetUserById(ctx context.Context, userId string) (u User, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "get", "name", u.Name, "userId", u.Id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.UserService.GetUserById(ctx, userId)
}

func (s *basicLoggingService) ListUsers(ctx context.Context) (us []User, err error) {
	defer func(begin time.Time) {
		s.logger.Log("operation", "list", "len", len(us), "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.UserService.ListUsers(ctx)
}

func NewBasicLoggingService(logger log.Logger, s Service) Service {
	return &basicLoggingService{
		logger:      logger,
		UserService: s,
	}
}
