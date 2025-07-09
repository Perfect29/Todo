package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Perfect29/Server/cmd/api/service"
	"github.com/Perfect29/Server/cmd/api/service/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoUseCase_GetByID(t *testing.T) {
	type arguments struct {
		ctx context.Context
		id int 
	}

	var tests = []struct {
		name string
		args arguments
		mockSetup func(repo *mock.TodoRepositoryMock, cache *mock.CacheMock, args arguments)
		wantResult *service.Todo
		wantErr bool
	} {
		{
			name: "Cache Hit",
			args: arguments{
				ctx: context.Background(),
				id: 1,
			},
			mockSetup: func(repo *mock.TodoRepositoryMock, cache *mock.CacheMock, args arguments) {
				cache.On("GetTodo", args.ctx, args.id).Return(&service.Todo{
					ID: 1,
					Name: "Cache hit",
					Description: "Cache hit description",
				}, nil)
			},
			wantResult: &service.Todo{
				ID: 1,
				Name: "Cache hit",
				Description: "Cache hit description",
			},
			wantErr: false,
		},

		{
			name: "Cache Miss, Repo Hit",
			args: arguments{
				ctx: context.Background(),
				id: 2,
			},
			mockSetup: func(repo *mock.TodoRepositoryMock, cache *mock.CacheMock, args arguments) {
				cache.On("GetTodo", args.ctx, args.id).Return(nil, nil)
				repo.On("GetByID", args.ctx, args.id).Return(&service.Todo{
					ID: 2,
				Name: "Cache Miss",
				Description: "Cache Miss description",
				}, nil)
				cache.On("SetTodo", args.ctx, &service.Todo{
					ID: 2,
					Name: "Cache Miss",
					Description: "Cache Miss description",
				}, 15*time.Minute).Return(nil)
			},
			wantResult: &service.Todo{
				ID: 2,
				Name: "Cache Miss",
				Description: "Cache Miss description",
			},
			wantErr: false,
		},

		{
			name: "cache miss and repo not found",
			args: arguments{
				ctx: context.Background(),
				id: 3,
			},
			mockSetup: func(repo *mock.TodoRepositoryMock, cache *mock.CacheMock, args arguments) {
				cache.On("GetTodo", args.ctx, args.id).Return(nil, nil)
				repo.On("GetByID", args.ctx, args.id).Return(nil, errors.New("Not found"))
			},
			wantResult: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t * testing.T) {
			t.Parallel()
			repo := new(mock.TodoRepositoryMock)
			cache := new(mock.CacheMock)

			tt.mockSetup(repo, cache, tt.args)

			u := &service.TodoUsecase{
				Repo: repo,
				Cache: cache,
			}

			todo, err := u.GetByID(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, todo)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantResult, todo)
			}

			repo.AssertExpectations(t)
			cache.AssertExpectations(t)
		})
	}
}

func TestTodoUseCase_AddTodo(t *testing.T) {
	type args struct {
		ctx context.Context
		todo *service.Todo
	}

	var tests = []struct {
		name string 
		args args
		mockSetup func(repo *mock.TodoRepositoryMock, args args) 
		wantErr bool
	} {
		{
			name: "Empty name",
			args: args{
				ctx: context.Background(),
				todo: &service.Todo{
					Name: "",
					Description: "",
					ID: 1,
				},
			},
			mockSetup: nil,
			wantErr: true,
		},

		{
			name: "Adding todo",
			args: args{
				ctx: context.Background(),
				todo: &service.Todo{
					Name: "Name",
					Description: "Description",
					ID: 2,
				},
			},
			mockSetup: func(repo *mock.TodoRepositoryMock, args args) {
				repo.On("AddTodo", args.ctx, args.todo).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t * testing.T) {
			t.Parallel()
			repo := new(mock.TodoRepositoryMock)
			u := &service.TodoUsecase{
				Repo: repo,
				Cache: nil,
			}

			if tt.mockSetup != nil {
				tt.mockSetup(repo, tt.args)
			}

			err := u.AddTodo(tt.args.ctx, tt.args.todo) 
			if tt.wantErr {
				require.Error(t, err)
				repo.AssertNotCalled(t, "AddTodo", tt.args.ctx, tt.args.todo)
			} else {
				require.NoError(t, err)
			}
			
			repo.AssertExpectations(t)
		})
	}
} 
