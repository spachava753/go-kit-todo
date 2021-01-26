package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/spachava/go-kit-todo/cmd"
	todotransport "github.com/spachava/go-kit-todo/pkg/todo/transport/gorilla"
	usertransport "github.com/spachava/go-kit-todo/pkg/user/transport/gorilla"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "8080"
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	todoService, userService := cmd.InitApp(logger)

	r := mux.NewRouter()
	usertransport.MakeHandler(userService, r)
	todotransport.MakeHandler(todoService, r)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("app exited with err: %s", err)
	}
}
