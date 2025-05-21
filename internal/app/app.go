package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/098765432m/monthly_planner_backend/internal/config"
	"github.com/098765432m/monthly_planner_backend/internal/database"
	month_handler "github.com/098765432m/monthly_planner_backend/internal/handler"
	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
	month_repository "github.com/098765432m/monthly_planner_backend/internal/repository/month"
	month_service "github.com/098765432m/monthly_planner_backend/internal/service/month"
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

	// Root
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the homepage!"))
	})

	// Repository
	dayRepo := day_repository.New(conn)
	monthRepo := month_repository.New(conn)

	// Service
	monthService := month_service.NewMonthService(monthRepo, dayRepo)

	// Handler
	monthHandler := month_handler.NewMonthHandler(monthService)

	//Chi Mount
	r.Route("/api", func(api chi.Router) {

		api.Mount("/months", monthHandler.Routes())

	})
	http.ListenAndServe(serverPortStr, r)

	return nil
}
