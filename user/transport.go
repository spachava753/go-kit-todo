package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func createUserRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return createUserRequest{}, err
	}
	return req, nil
}

func createUserResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(createUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func deleteUserRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		return deleteUserRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return deleteUserRequest{userId}, nil
}

func deleteUserResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(deleteUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func updateUserRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return updateUserRequest{}, err
	}
	return updateUserRequest{user}, nil
}

func updateUserResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(updateUserResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func getUserByIdRequestDecoder(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		return getUserByIdRequest{}, fmt.Errorf("user id is not provided in path")
	}
	return getUserByIdRequest{userId}, nil
}

func getUserByIdResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(getUserByIdResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func listUsersRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return listUsersRequest{}, err
}

func listUsersResponseEncoder(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(listUsersResponse)
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

func MakeHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	createUserHandler := httptransport.NewServer(
		makeHttpCreateUserEndpoint(s),
		createUserRequestDecoder,
		createUserResponseEncoder,
	)
	r.Handle("/user", createUserHandler).Methods("PUT")

	deleteUserHandler := httptransport.NewServer(
		makeHttpDeleteUserEndpoint(s),
		deleteUserRequestDecoder,
		deleteUserResponseEncoder,
	)
	r.Handle("/user/{id}", deleteUserHandler).Methods("DELETE")

	updateUserHandler := httptransport.NewServer(
		makeHttpUpdateUserEndpoint(s),
		updateUserRequestDecoder,
		updateUserResponseEncoder,
	)
	r.Handle("/user", updateUserHandler).Methods("POST")

	getUserHandler := httptransport.NewServer(
		makeHttpGetUserByIdEndpoint(s),
		getUserByIdRequestDecoder,
		getUserByIdResponseEncoder,
	)
	r.Handle("/user/{id}", getUserHandler).Methods("GET")

	listUserHandler := httptransport.NewServer(
		makeHttpListUsersEndpoint(s),
		listUsersRequestDecoder,
		listUsersResponseEncoder,
	)
	r.Handle("/user", listUserHandler).Methods("GET")

	return r
}
