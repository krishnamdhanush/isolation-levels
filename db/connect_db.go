package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
