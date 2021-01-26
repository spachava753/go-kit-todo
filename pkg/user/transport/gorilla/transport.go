package gorilla

import (
	"context"
	"encoding/json"
	"fmt"
	kittransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/spachava/go-kit-todo/pkg/user"
	"net/http"
)

func createUserRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return user.CreateUserRequest{}, err
	}
	return req, nil
}

func createUserResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(user.CreateUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func deleteUserRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		return user.DeleteUserRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return user.DeleteUserRequest{UserId: userId}, nil
}

func deleteUserResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(user.DeleteUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func updateUserRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return user.UpdateUserRequest{}, err
	}
	return user.UpdateUserRequest{User: u}, nil
}

func updateUserResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(user.UpdateUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func getUserByIdRequestDecoder(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		return user.GetUserByIdRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return user.GetUserByIdRequest{UserId: userId}, nil
}

func getUserByIdResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(user.GetUserByIdResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func listUsersRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return user.ListUsersRequest{}, err
}

func listUsersResponseEncoder(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(user.ListUsersResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func MakeHandler(s user.Service, r *mux.Router) {
	createUserHandler := kittransport.NewServer(
		user.MakeCreateUserEndpoint(s),
		createUserRequestDecoder,
		createUserResponseEncoder,
	)
	r.Handle("/user", createUserHandler).Methods("PUT")

	deleteUserHandler := kittransport.NewServer(
		user.MakeDeleteUserEndpoint(s),
		deleteUserRequestDecoder,
		deleteUserResponseEncoder,
	)
	r.Handle("/user/{id}", deleteUserHandler).Methods("DELETE")

	updateUserHandler := kittransport.NewServer(
		user.MakeUpdateUserEndpoint(s),
		updateUserRequestDecoder,
		updateUserResponseEncoder,
	)
	r.Handle("/user", updateUserHandler).Methods("POST")

	getUserHandler := kittransport.NewServer(
		user.MakeGetUserByIdEndpoint(s),
		getUserByIdRequestDecoder,
		getUserByIdResponseEncoder,
	)
	r.Handle("/user/{id}", getUserHandler).Methods("GET")

	listUserHandler := kittransport.NewServer(
		user.MakeListUsersEndpoint(s),
		listUsersRequestDecoder,
		listUsersResponseEncoder,
	)
	r.Handle("/user", listUserHandler).Methods("GET")
}
