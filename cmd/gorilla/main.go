package main

import (
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
	"github.com/spachava/go-kit-todo/todo"
	todotransport "github.com/spachava/go-kit-todo/todo/transport/gorilla"
	"github.com/spachava/go-kit-todo/user"
	usertransport "github.com/spachava/go-kit-todo/user/transport/gorilla"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {

	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
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

	r := mux.NewRouter()
	usertransport.MakeHandler(userService, r)
	todotransport.MakeHandler(todoService, r)
	http.ListenAndServe(":8080", r)
}
