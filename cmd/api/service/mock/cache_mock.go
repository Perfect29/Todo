package mock

import (
	"context"
	"time"

	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/stretchr/testify/mock"
)

type CacheMock struct {
	mock.Mock
}

func (m * CacheMock) GetTodo(ctx context.Context, id int) (*service.Todo, error) {
	args := m.Called(ctx, id)

	todo, ok := args.Get(0).(*service.Todo) 
	if !ok && args.Get(0) != nil {
		return nil, args.Error(1)
	}
	return todo, args.Error(1)
}

func (m *CacheMock) SetTodo(ctx context.Context, todo *service.Todo, ttl time.Duration) error {
	args := m.Called(ctx, todo, ttl)

	return args.Error(0)
}

func (m *CacheMock) DeleteTodo(ctx context.Context, id int) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

