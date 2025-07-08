// package service

// import (
// 	"context"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/google/uuid"
// )

// type dbService struct {
// 	db *pgx.Conn
// }

// func InitDB(ctx context.Context) (*pgx.Conn, error){
// 	DB, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")

// 	if err != nil {
// 		return nil, err
// 	}
// 	return DB, nil
// }

// func (DB *dbService) AddTodo(todo *Todo) error{
// 	query := `
// 		INSERT INTO todo (name, description)
// 		VALUES ($1, $2)
// 		RETURNING id
// 	`
// 	todo.ID = uuid.NewString()
// }