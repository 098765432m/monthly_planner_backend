package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/098765432m/monthly_planner_backend/internal/config"
	"github.com/098765432m/monthly_planner_backend/internal/database"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Run() error {
	// Init Global logger
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))

	// Init Config environment file
	if err := config.LoadConfig(); err != nil {
		return err
	}

	//Connect to Database
	conn, err := database.Connect() // PostgreSQL
	if err != nil {
		zap.S().Fatal("Failed to connect to PostgreSQL")
	}

	defer conn.Close(context.Background())

	//Initialize server
	r := chi.NewRouter()
	serverPort := config.AppGlobalConfigData.App.Port
	serverPortStr := fmt.Sprintf(":%s", serverPort)

	zap.S().Infof("Server running at port %s!", serverPort)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the homepage!"))
	})

	http.ListenAndServe(serverPortStr, r)

	return nil
}
