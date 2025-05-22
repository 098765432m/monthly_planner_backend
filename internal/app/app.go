package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/098765432m/monthly_planner_backend/internal/config"
	"github.com/098765432m/monthly_planner_backend/internal/database"
	month_handler "github.com/098765432m/monthly_planner_backend/internal/handler/month"
	user_handler "github.com/098765432m/monthly_planner_backend/internal/handler/user"

	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
	month_repository "github.com/098765432m/monthly_planner_backend/internal/repository/month"
	user_repository "github.com/098765432m/monthly_planner_backend/internal/repository/user"
	month_service "github.com/098765432m/monthly_planner_backend/internal/service/month"
	user_service "github.com/098765432m/monthly_planner_backend/internal/service/user"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Run() error {

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
	userRepo := user_repository.New(conn)
	dayRepo := day_repository.New(conn)
	monthRepo := month_repository.New(conn)

	// Service
	userService := user_service.NewUserService(userRepo)
	monthService := month_service.NewMonthService(monthRepo, dayRepo)

	// Handler
	userHandler := user_handler.NewUserHandler(userService)
	monthHandler := month_handler.NewMonthHandler(monthService)

	//Chi Mount
	r.Route("/api", func(api chi.Router) {

		api.Mount("/users", userHandler.RegisterRoutes())
		api.Mount("/months", monthHandler.RegisterRoutes())

	})
	http.ListenAndServe(serverPortStr, r)

	return nil
}
