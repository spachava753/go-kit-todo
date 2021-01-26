package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spachava/go-kit-todo/cmd"
	todotransport "github.com/spachava/go-kit-todo/pkg/todo/transport/fiber"
	usertransport "github.com/spachava/go-kit-todo/pkg/user/transport/fiber"
	"os"
)

func main() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "8080"
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	todoService, userService := cmd.InitApp(logger)

	app := fiber.New()

	app.Use(recover.New())
	todotransport.MakeRoutes(todoService, app)
	usertransport.MakeRoutes(userService, app)

	if err := app.Listen(":" + port); err != nil {
		fmt.Printf("app exited with err: %s", err)
	}
}
