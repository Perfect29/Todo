package service

import (
	log "github.com/sirupsen/logrus"
	"context"
	"github.com/jackc/pgx/v5"
)

func NewPostgresRepository(db *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func InitDB(ctx context.Context) (*PostgresRepository, error){
	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@db:5432/postgres")

	if err != nil {
		return nil, err
	}
	log.Info("Database Repository was successfully created")
	return NewPostgresRepository(conn), nil
}