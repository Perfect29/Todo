package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TodoRepository interface {
	AddTodo(ctx context.Context, todo *Todo) error
	RemoveTodo(ctx context.Context, id int) error
	ShowListTodo(ctx context.Context) (ToDoList, error)
	GetByID(ctx context.Context, id int) (*Todo, error)
}

type PostgresRepository struct {
	db *pgx.Conn
}

func (db *PostgresRepository) AddTodo(ctx context.Context, todo *Todo) error{
	query := `
		INSERT INTO todo (name, description)
		VALUES ($1, $2)
		RETURNING id
	`
	err := db.db.QueryRow(ctx, query, todo.Name, todo.Description).Scan(&todo.ID)
	if err != nil {
		return err
	}
	return nil
}


func (db *PostgresRepository) RemoveTodo(ctx context.Context, id int) error {
	query := `
		DELETE FROM todo
		WHERE
		id = $1
	`
	result, err := db.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo with ID %d was not found", id)
	}
	return nil
}

func (db * PostgresRepository) ShowListTodo(ctx context.Context) (ToDoList, error) {
	query := `
		SELECT id, name, description FROM todo
	`
	rows, err := db.db.Query(ctx, query)
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

func (db *PostgresRepository) GetByID(ctx context.Context, id int) (*Todo, error) {
	query := `
		SELECT name, description
		FROM todo
		WHERE id = $1
	`
	row := db.db.QueryRow(ctx, query, id)
	var todo Todo
	todo.ID = id
	err := row.Scan(&todo.Name, &todo.Description)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (db *PostgresRepository) Close(ctx context.Context) error {
	return db.db.Close(ctx)
}