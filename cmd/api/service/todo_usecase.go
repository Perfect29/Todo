package service

import (
	"context"
	"fmt"
	"time"

)

type TodoUsecase struct {
	Repo TodoRepository
	Cache Cache
}

func (u *TodoUsecase) validateTodo(todo *Todo) error {
	if todo.Name == "" {
		return fmt.Errorf("todo name is required")
	}
	return nil
}

func (u *TodoUsecase) GetByID(ctx context.Context, id int) (*Todo, error) {
	todo, err := u.Cache.GetTodo(ctx, id)
	if err != nil {
		fmt.Println("cache error: ", err)
	}
	if todo != nil {
		return todo, nil
	}
	todo, err = u.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = u.Cache.SetTodo(ctx, todo, 15 * time.Minute); err != nil {
		fmt.Println("falied to set cache: ", err)
	}
	return todo, nil
}

func (u *TodoUsecase) AddTodo(ctx context.Context, todo *Todo) error {
	if err := u.validateTodo(todo); err != nil {
		return err
	}
	return u.Repo.AddTodo(ctx, todo)
}

func (u *TodoUsecase) ShowListTodo(ctx context.Context) (ToDoList, error) {
	return u.Repo.ShowListTodo(ctx)
}

func (u *TodoUsecase) RemoveTodo(ctx context.Context, id int) error {
	err := u.Repo.RemoveTodo(ctx, id)
	if err != nil {
		return err
	}
	if err = u.Cache.DeleteTodo(ctx, id); err != nil {
		fmt.Println("failed to delete from cache: ", err)
	}
	return nil
}