package main

import (
	"context"
	"os"

	db "autoerp/db/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func connectDB() (*db.Queries, *pgx.Conn, error) {

	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		return nil, nil, err
	}

	queries := db.New(conn)

	return queries, conn, nil
}