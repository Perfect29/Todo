package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	DB *pgx.Conn
	Cache *CacheService
}

func InitDB(ctx context.Context) (*Service, error){
	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@db:5432/postgres")

	if err != nil {
		return nil, err
	}
	return &Service{DB: conn}, nil
}

func (s *Service) AddTodo(ctx context.Context, todo *Todo) error{
	query := `
		INSERT INTO todo (name, description)
		VALUES ($1, $2)
		RETURNING id
	`
	err := s.DB.QueryRow(ctx, query, todo.Name, todo.Description).Scan(&todo.ID)
	if err != nil {
		return err
	}
	return nil
}


func (s *Service) RemoveTodo(ctx context.Context, id int) error {
	query := `
		DELETE FROM todo
		WHERE
		id = $1
	`
	result, err := s.DB.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo with ID %d was not found", id)
	}
	s.Cache.DeleteTodo(id)
	return nil
}

func (s * Service) ShowListTodo(ctx context.Context) (ToDoList, error) {
	query := `
		SELECT id, name, description FROM todo
	`
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo 
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*Todo, error) {
	toDo, err := s.Cache.GetTodo(id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if toDo != nil {
		return toDo, nil 
	}

	query := `
		SELECT name, description
		FROM todo
		WHERE id = $1
	`
	row := s.DB.QueryRow(ctx, query, id)
	var todo Todo
	todo.ID = id
	err = row.Scan(&todo.Name, &todo.Description)
	if err != nil {
		return nil, err
	}
	s.Cache.SetTodo(&todo, 15 * time.Minute)
	return &todo, nil
}

func (s *Service) Close(ctx context.Context) error {
	return s.DB.Close(ctx)
}