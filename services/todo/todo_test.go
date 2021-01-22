package todo

import (
	"context"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewMemTodoService(t *testing.T) {
	tm := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(tm.UnixNano())), 0)
	svc := NewMemTodoService(entropy, tm)
	if svc == nil {
		t.Errorf("expected a non-nil object, got nil")
	}
}

func Test_inMemTodoService_CreateTodo(t *testing.T) {
	userId := "abc123"
	type args struct {
		ctx  context.Context
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Simple Todo",
			args: args{
				ctx:  nil,
				text: "Take out trash",
			},
			wantErr: false,
		},
		{
			name: "Simple Todo",
			args: args{
				ctx:  nil,
				text: "Sleep early",
			},
			wantErr: false,
		},
		{
			name: "Invalid Todo",
			args: args{
				ctx:  nil,
				text: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := time.Unix(1000000, 0)
			entropy := ulid.Monotonic(rand.New(rand.NewSource(tm.UnixNano())), 0)
			s := NewMemTodoService(entropy, tm)
			got, err := s.CreateTodo(tt.args.ctx, tt.args.text, userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.UserId == "" && !tt.wantErr {
				t.Errorf("CreateTodo() UserId is empty")
				return
			}
			if got.Id == "" && !tt.wantErr {
				t.Errorf("CreateTodo() Id is empty")
				return
			}
			if got.Text != tt.args.text {
				t.Errorf("CreateTodo() Text = %v, want %v", err, tt.args.text)
				return
			}
		})
	}
}

func Test_inMemTodoService_DeleteTodo(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		todos   sync.Map
	}
	type args struct {
		ctx    context.Context
		todoId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemTodoService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				todos:   tt.fields.todos,
			}
			got, err := s.DeleteTodo(tt.args.ctx, tt.args.todoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemTodoService_GetTodoById(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		todos   sync.Map
	}
	type args struct {
		ctx    context.Context
		todoId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemTodoService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				todos:   tt.fields.todos,
			}
			got, err := s.GetTodoById(tt.args.ctx, tt.args.todoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodoById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemTodoService_ListTodosByUserId(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		todos   sync.Map
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemTodoService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				todos:   tt.fields.todos,
			}
			got, err := s.ListTodosByUserId(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTodosByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTodosByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemTodoService_UpdateTodo(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		todos   sync.Map
	}
	type args struct {
		ctx  context.Context
		todo Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemTodoService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				todos:   tt.fields.todos,
			}
			got, err := s.UpdateTodo(tt.args.ctx, tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
