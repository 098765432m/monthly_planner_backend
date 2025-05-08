package main

import (
	"github.com/098765432m/monthly_planner_backend/internal/app"
	"go.uber.org/zap"
)

func main() {
	if err := app.Run(); err != nil {
		zap.L().Sugar().Fatal(err)
	}
}
