package fiber

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/spachava/go-kit-todo/pkg/todo"
	"github.com/spachava/go-kit-todo/pkg/transport"
)

func createTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	var req todo.CreateTodoRequest
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return todo.CreateTodoRequest{}, err
	}
	return req, nil
}

func deleteTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	todoId := ctx.Params("id", "")
	if todoId == "" {
		return todo.DeleteTodoRequest{}, errors.New("missing todo id in path")
	}
	return todo.DeleteTodoRequest{TodoId: todoId}, nil
}

func updateTodoRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	var t todo.Todo
	if err := json.Unmarshal(ctx.Body(), &t); err != nil {
		return todo.UpdateTodoRequest{}, err
	}
	return todo.UpdateTodoRequest{Todo: t}, nil
}

func getTodoByIdRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	todoId := ctx.Params("id", "")
	if todoId == "" {
		return todo.GetTodoByIdRequest{}, errors.New("missing todo id in path")
	}
	return todo.GetTodoByIdRequest{TodoId: todoId}, nil
}

func getTodosByUserIdRequestEncoder(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("userId", "")
	if userId == "" {
		return todo.GetTodosByUserIdRequest{}, errors.New("missing user id in path")
	}
	return todo.GetTodosByUserIdRequest{UserId: userId}, nil
}

func makeResponseDecoder() func(ctx *fiber.Ctx, resp interface{}) error {
	return func(ctx *fiber.Ctx, resp interface{}) error {
		if err := ctx.JSON(resp); err != nil {
			return err
		}
		return nil
	}
}

func MakeRoutes(s todo.Service, app *fiber.App) {
	createTodoHandler := transport.MakeFiberHandler(
		todo.MakeCreateTodoEndpoint(s),
		createTodoRequestEncoder,
		makeResponseDecoder(),
		nil,
	)
	app.Put("/todo", createTodoHandler)

	deleteTodoHandler := transport.MakeFiberHandler(
		todo.MakeDeleteTodoEndpoint(s),
		deleteTodoRequestEncoder,
		makeResponseDecoder(),
		nil,
	)
	app.Delete("/todo/:id", deleteTodoHandler)

	updateTodoHandler := transport.MakeFiberHandler(
		todo.MakeUpdateTodoEndpoint(s),
		updateTodoRequestEncoder,
		makeResponseDecoder(),
		nil,
	)
	app.Post("/todo", updateTodoHandler)

	getTodoByIdHandler := transport.MakeFiberHandler(
		todo.MakeGetTodoByIdEndpoint(s),
		getTodoByIdRequestEncoder,
		makeResponseDecoder(),
		nil,
	)
	app.Get("/todo/:id", getTodoByIdHandler)

	getTodoByUserIdTodoHandler := transport.MakeFiberHandler(
		todo.MakeGetTodosByUserIdEndpoint(s),
		getTodosByUserIdRequestEncoder,
		makeResponseDecoder(),
		nil,
	)
	app.Get("/todo/list/:userId", getTodoByUserIdTodoHandler)
}
