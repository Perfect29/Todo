package service

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type TodoUsecase struct {
	Repo TodoRepository
	Cache Cache
}

func (u *TodoUsecase) validateTodo(todo *Todo) error {
	if todo.Name == "" {
		log.Error("Validation error: todo name is empty")
		return fmt.Errorf("todo name is required")
	}
	return nil
}

func (u *TodoUsecase) GetByID(ctx context.Context, id int) (*Todo, error) {
	todo, err := u.Cache.GetTodo(ctx, id)
	if err != nil {
		log.Warnf("Cache GetTodo failed for id %d: %v", id, err)
	}
	if todo != nil {
		log.Info("Cache Hit")
		return todo, nil
	}
	todo, err = u.Repo.GetByID(ctx, id)
	if err != nil {
		log.Warnf("No todo with id = %d", id)
		return nil, err
	}
	if err = u.Cache.SetTodo(ctx, todo, 15 * time.Minute); err != nil {
		log.Warnf("failed set todo with id %d to cache: %v", id, err)
	}
	log.Infof("todo with id %d retrieved successfully from database", id)
	return todo, nil
}

func (u *TodoUsecase) AddTodo(ctx context.Context, todo *Todo) error {
	if err := u.validateTodo(todo); err != nil {
		return err
	}
	if err := u.Repo.AddTodo(ctx, todo); err != nil {
		log.Errorf("failed to add todo  %v", err)
		return err
	}
	log.Info("new todo was added successfully")
	return nil
}

func (u *TodoUsecase) ShowListTodo(ctx context.Context) (ToDoList, error) {
	list, err := u.Repo.ShowListTodo(ctx)
	if err != nil {
		log.Error("Failed to retrieve showlist from database", err)
		return nil, err
	}
	log.Infof("Retreived showlist with %d todos", len(list))
	return list, nil
}

func (u *TodoUsecase) RemoveTodo(ctx context.Context, id int) error {
	err := u.Repo.RemoveTodo(ctx, id)
	if err != nil {
		log.Errorf("Could not remove todo with id %d: %v", id, err)
		return err
	}
	if err = u.Cache.DeleteTodo(ctx, id); err != nil {
		log.Warn("Failed to delete from cache:", err)
	}
	return nil
}