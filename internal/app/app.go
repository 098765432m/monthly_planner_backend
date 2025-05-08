package app

import (
	"github.com/098765432m/monthly_planner_backend/internal/config"
	"go.uber.org/zap"
)

func Run() error {
	// Init Global logger
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))

	// Init Config environment file
	if err := config.InitConfig(); err != nil {
		return err
	}

	dbPort := config.AppGlobalConfigData.Database.Port

	zap.S().Info(dbPort)

	return nil
}
