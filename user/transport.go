package user

import (
	"context"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func createUserRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return nil, err
}

func createUserResponseEncoder(context.Context, http.ResponseWriter, interface{}) error {
	return nil
}

func deleteUserRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return nil, err
}

func deleteUserResponseEncoder(context.Context, http.ResponseWriter, interface{}) error {
	return nil
}

func updateUserRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return nil, err
}

func updateUserResponseEncoder(context.Context, http.ResponseWriter, interface{}) error {
	return nil
}

func getUserByIdRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return nil, err
}

func getUserByIdResponseEncoder(context.Context, http.ResponseWriter, interface{}) error {
	return nil
}

func listUsersRequestDecoder(context.Context, *http.Request) (request interface{}, err error) {
	return nil, err
}

func listUsersResponseEncoder(context.Context, http.ResponseWriter, interface{}) error {
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
