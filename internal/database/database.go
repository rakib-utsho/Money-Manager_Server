package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		return err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return err
	}

	DB = db

	fmt.Println("Database connected successfully")

	return nil
}
