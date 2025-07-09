package mock

import (
	"context"

	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct{
	mock.Mock
}

func (m * TodoRepositoryMock) AddTodo(ctx context.Context, todo *service.Todo) error {
	args := m.Called(ctx, todo)

	return args.Error(0)
}

func (m * TodoRepositoryMock) RemoveTodo(ctx context.Context, id int) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m * TodoRepositoryMock) ShowListTodo(ctx context.Context) (service.ToDoList, error) {
	args := m.Called(ctx)

	todos, ok := args.Get(0).(service.ToDoList)

	if !ok && args.Get(0) != nil {
		return nil, args.Error(1)
	}

	return todos, args.Error(1)
}

func (m * TodoRepositoryMock) GetByID(ctx context.Context, id int) (*service.Todo, error) {
	args := m.Called(ctx, id)

	todo, ok := args.Get(0).(*service.Todo)

	if !ok && args.Get(0) != nil {
		return nil, args.Error(1)
	}
	
	return todo, args.Error(1)
}	