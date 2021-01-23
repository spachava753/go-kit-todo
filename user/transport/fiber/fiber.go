package user

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spachava/go-kit-todo/transport"
	"github.com/spachava/go-kit-todo/user"
)

func createUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	ctx.Accepts("json", "application/json")
	var req user.CreateUserRequest
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return user.CreateUserRequest{}, err
	}
	return req, nil
}

func createUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(user.CreateUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func deleteUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("id", "")
	if userId == "" {
		return user.DeleteUserRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return user.DeleteUserRequest{userId}, nil
}

func deleteUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(user.DeleteUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func updateUserRequestDecoder(ctx *fiber.Ctx) (interface{}, error) {
	var u user.User
	if err := json.Unmarshal(ctx.Body(), &u); err != nil {
		return user.UpdateUserRequest{}, err
	}
	return user.UpdateUserRequest{u}, nil
}

func updateUserResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(user.UpdateUserResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func getUserByIdRequestDecoder(ctx *fiber.Ctx) (request interface{}, err error) {
	userId := ctx.Params("id", "")
	if userId == "" {
		return user.GetUserByIdRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return user.GetUserByIdRequest{userId}, nil
}

func getUserByIdResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(user.GetUserByIdResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func listUsersRequestDecoder(ctx *fiber.Ctx) (request interface{}, err error) {
	return user.ListUsersRequest{}, err
}

func listUsersResponseEncoder(ctx *fiber.Ctx, resp interface{}) error {
	r := resp.(user.ListUsersResponse)
	if err := ctx.JSON(r); err != nil {
		return err
	}
	return nil
}

func MakeRoutes(s user.Service, app *fiber.App) {
	createUserHandler := transport.MakeFiberHandler(
		user.MakeCreateUserEndpoint(s),
		createUserRequestDecoder,
		createUserResponseEncoder,
	)
	app.Put("/user", createUserHandler)

	deleteUserHandler := transport.MakeFiberHandler(
		user.MakeDeleteUserEndpoint(s),
		deleteUserRequestDecoder,
		deleteUserResponseEncoder,
	)
	app.Delete("/user/:id", deleteUserHandler)

	updateUserHandler := transport.MakeFiberHandler(
		user.MakeUpdateUserEndpoint(s),
		updateUserRequestDecoder,
		updateUserResponseEncoder,
	)
	app.Post("/user", updateUserHandler)

	getUserHandler := transport.MakeFiberHandler(
		user.MakeGetUserByIdEndpoint(s),
		getUserByIdRequestDecoder,
		getUserByIdResponseEncoder,
	)
	app.Get("/user/:id", getUserHandler)

	listUserHandler := transport.MakeFiberHandler(
		user.MakeListUsersEndpoint(s),
		listUsersRequestDecoder,
		listUsersResponseEncoder,
	)
	app.Get("/user", listUserHandler)
}
