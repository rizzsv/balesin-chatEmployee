package database

import (
	"context"
	"os"
	"time"

	"balesin-chatEmployee/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to ping database")
	}

	DB = pool
}
