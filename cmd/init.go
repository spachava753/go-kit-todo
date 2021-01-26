package cmd

import (
	"github.com/go-kit/kit/log"
	"github.com/oklog/ulid/v2"
	"github.com/spachava/go-kit-todo/pkg/todo"
	"github.com/spachava/go-kit-todo/pkg/user"
	"math/rand"
	"time"
)

func InitApp(logger log.Logger) (todo.Service, user.Service) {
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)

	t := time.Unix(1000000, 0)
	userEntropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	var userService user.Service
	userService = user.NewMemUserService(userEntropy, t)
	userService = user.NewBasicLoggingService(log.With(logger, "component", "user"), userService)

	todoEntropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	var todoService todo.Service
	todoService = todo.NewMemTodoService(todoEntropy, t)
	todoService = todo.NewBasicLoggingService(log.With(logger, "component", "todo"), todoService)

	return todoService, userService
}
