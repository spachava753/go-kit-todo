package gorilla

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/spachava/go-kit-todo/todo"
	"net/http"
)

func createTodoRequestEncoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req todo.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return todo.CreateTodoRequest{}, err
	}
	return req, nil
}

func deleteTodoRequestEncoder(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	todoId, ok := vars["id"]
	if !ok {
		return todo.DeleteTodoRequest{}, errors.New("missing todo id in path")
	}
	return todo.DeleteTodoRequest{TodoId: todoId}, nil
}

func updateTodoRequestEncoder(_ context.Context, r *http.Request) (interface{}, error) {
	var t todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return todo.UpdateTodoRequest{}, err
	}
	return todo.UpdateTodoRequest{Todo: t}, nil
}

func getTodoByIdRequestEncoder(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	todoId, ok := vars["id"]
	if !ok {
		return todo.GetTodoByIdRequest{}, errors.New("missing todo id in path")
	}
	return todo.GetTodoByIdRequest{TodoId: todoId}, nil
}

func getTodosByUserIdRequestEncoder(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return todo.GetTodosByUserIdRequest{}, errors.New("missing user id in path")
	}
	return todo.GetTodosByUserIdRequest{UserId: userId}, nil
}

func makeResponseDecoder() func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			return err
		}
		return nil
	}
}

func MakeHandler(s todo.Service, r *mux.Router) {
	createTodoHandler := httptransport.NewServer(
		todo.MakeCreateTodoEndpoint(s),
		createTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo", createTodoHandler).Methods("PUT")

	deleteTodoHandler := httptransport.NewServer(
		todo.MakeDeleteTodoEndpoint(s),
		deleteTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/{id}", deleteTodoHandler).Methods("DELETE")

	updateTodoHandler := httptransport.NewServer(
		todo.MakeUpdateTodoEndpoint(s),
		updateTodoRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo", updateTodoHandler).Methods("POST")

	getTodoByIdHandler := httptransport.NewServer(
		todo.MakeGetTodoByIdEndpoint(s),
		getTodoByIdRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/{id}", getTodoByIdHandler).Methods("GET")

	getTodoByUserIdTodoHandler := httptransport.NewServer(
		todo.MakeListTodosByUserIdEndpoint(s),
		getTodosByUserIdRequestEncoder,
		makeResponseDecoder(),
	)
	r.Handle("/todo/list/{userId}", getTodoByUserIdTodoHandler).Methods("GET")
}
