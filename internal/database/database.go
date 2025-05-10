package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Connect() (*pgx.Conn, error) {
	// set timeout for db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Create dsn string
	connectStr := viper.GetString("DB_URL")
	if connectStr == "" {
		zap.S().Fatal("DB_URL is not set!")
	}

	// Connect to Database
	conn, err := pgx.Connect(ctx, connectStr)
	if err != nil {
		zap.S().Errorln(err)
		zap.S().Fatalln("Failed to connect to database!")
	}

	zap.S().Infoln("Connected to PostgreSQL successfully!")

	return conn, nil
}
