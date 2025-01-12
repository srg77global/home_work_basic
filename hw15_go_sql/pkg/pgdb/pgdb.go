package pgdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func New(dbDSN string, maxOpenConns int32) (context.Context, *pgxpool.Pool, error) {
	ctx := context.Background()

	connConfig, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		return ctx, nil, fmt.Errorf("error creating DSN: %w", err)
	}

	connConfig.MaxConns = maxOpenConns
	connConfig.MinConns = 0

	dbc, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return ctx, nil, fmt.Errorf("error connecting DB: %w", err)
	}

	if err = dbc.Ping(ctx); err != nil {
		return ctx, nil, fmt.Errorf("error ping DB: %w", err)
	}

	return ctx, dbc, nil
}

func LoadENV() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
		return "", err
	}
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if !exists {
		log.Println("DSN: ", dbDSN)
		return "", errors.New("no DSN variable in .env found")
	}
	return dbDSN, nil
}
