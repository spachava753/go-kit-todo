package user

import (
	"context"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewMemUserService(t *testing.T) {
	tm := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(tm.UnixNano())), 0)
	svc := NewMemUserService(entropy, tm)
	if svc == nil {
		t.Errorf("expected a non-nil object, got nil")
	}
}

func Test_inMemUserService_CreateUser(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		users   sync.Map
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemUserService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				users:   tt.fields.users,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemUserService_DeleteUser(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		users   sync.Map
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemUserService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				users:   tt.fields.users,
			}
			got, err := s.DeleteUser(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemUserService_GetUserById(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		users   sync.Map
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemUserService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				users:   tt.fields.users,
			}
			got, err := s.GetUserById(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemUserService_ListUsers(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		users   sync.Map
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemUserService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				users:   tt.fields.users,
			}
			got, err := s.ListUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemUserService_UpdateUser(t *testing.T) {
	type fields struct {
		entropy *ulid.MonotonicEntropy
		time    time.Time
		users   sync.Map
	}
	type args struct {
		ctx  context.Context
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemUserService{
				entropy: tt.fields.entropy,
				time:    tt.fields.time,
				users:   tt.fields.users,
			}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
