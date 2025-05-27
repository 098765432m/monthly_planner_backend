package app

import (
	"context"
	"fmt"

	"github.com/098765432m/monthly_planner_backend/internal/config"
	"github.com/098765432m/monthly_planner_backend/internal/database"
	month_handler "github.com/098765432m/monthly_planner_backend/internal/handler/month"
	task_category_handler "github.com/098765432m/monthly_planner_backend/internal/handler/task_category"
	user_handler "github.com/098765432m/monthly_planner_backend/internal/handler/user"
	"github.com/gin-gonic/gin"

	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
	month_repository "github.com/098765432m/monthly_planner_backend/internal/repository/month"
	task_repository "github.com/098765432m/monthly_planner_backend/internal/repository/task"
	task_category_repository "github.com/098765432m/monthly_planner_backend/internal/repository/task_category"
	user_repository "github.com/098765432m/monthly_planner_backend/internal/repository/user"
	month_service "github.com/098765432m/monthly_planner_backend/internal/service/month"
	task_category_service "github.com/098765432m/monthly_planner_backend/internal/service/task_category"
	user_service "github.com/098765432m/monthly_planner_backend/internal/service/user"
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
	r := gin.Default()
	// r := chi.NewRouter()
	serverPort := config.AppGlobalConfigData.App.Port
	serverPortStr := fmt.Sprintf(":%s", serverPort)

	zap.S().Infof("Server running at port %s!", serverPort)

	// Repository
	userRepo := user_repository.New(conn)
	dayRepo := day_repository.New(conn)
	monthRepo := month_repository.New(conn)
	taskCategoryRepo := task_category_repository.New(conn)
	taskRepo := task_repository.New(conn)

	// Service
	userService := user_service.NewUserService(userRepo)
	monthService := month_service.NewMonthService(
		&month_service.MonthServiceDesp{
			Repo:      monthRepo,
			Day_repo:  dayRepo,
			Task_Repo: taskRepo,
		},
	)
	taskCategoryService := task_category_service.NewTaskCategoryService(taskCategoryRepo)

	// Handler
	userHandler := user_handler.NewUserHandler(userService)
	monthHandler := month_handler.NewMonthHandler(monthService)
	taskCategoryHandler := task_category_handler.NewTaskCategoryHandler(taskCategoryService)

	//Gin Group
	api := r.Group("/api")
	userHandler.RegisterRoutes(api)
	monthHandler.RegisterRoutes(api)
	taskCategoryHandler.RegisterRoutes(api)

	return r.Run(serverPortStr)
}
