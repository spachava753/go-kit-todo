package todo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func createTodoRequestEncoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return createTodoRequest{}, err
	}
	return req, nil
}

func deleteTodoRequestEncoder(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	todoId, ok := vars["id"]
	if !ok {
		return deleteTodoRequest{}, errors.New("missing todo id in path")
	}
	return deleteTodoRequest{todoId}, nil
}

func updateTodoRequestEncoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return updateTodoRequest{}, err
	}
	return updateTodoRequest{todo}, nil
}

func getTodoByIdRequestEncoder(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	todoId, ok := vars["id"]
	if !ok {
		return getTodoByIdRequest{}, errors.New("missing todo id in path")
	}
	return getTodoByIdRequest{todoId}, nil
}

func getTodosByUserIdRequestEncoder(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return getTodosByUserIdRequest{}, errors.New("missing user id in path")
	}
	return getTodosByUserIdRequest{userId}, nil
}

func makeResponseDecoder() func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			return err
		}
		return nil
	}
}

func MakeHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	createTodoHandler := httptransport.NewServer(
		makeHttpCreateTodoEndpoint(s),
		createTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo", createTodoHandler).Methods("PUT")

	deleteTodoHandler := httptransport.NewServer(
		makeHttpDeleteTodoEndpoint(s),
		deleteTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/{id}", deleteTodoHandler).Methods("DELETE")

	updateTodoHandler := httptransport.NewServer(
		makeHttpUpdateTodoEndpoint(s),
		updateTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo", updateTodoHandler).Methods("POST")

	getTodoByIdHandler := httptransport.NewServer(
		makeHttpGetTodoByIdEndpoint(s),
		getTodoByIdRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/{id}", getTodoByIdHandler).Methods("GET")

	getTodoByUserIdTodoHandler := httptransport.NewServer(
		makeHttpListTodosByUserIdEndpoint(s),
		getTodosByUserIdRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/list/{userId}", getTodoByUserIdTodoHandler).Methods("GET")

	return r
}
