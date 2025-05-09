package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Connect() *sql.DB {

	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)

	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		zap.S().Errorln(err)
		zap.S().Fatalln("Failed to connect to database!")
	}

	if err = db.Ping(); err != nil {
		zap.S().Errorln(err)
		zap.S().Fatalln("DB is not reachable!")
	}

	zap.S().Infoln("Connected to PostgreSQL successfully!")

	return db
}
