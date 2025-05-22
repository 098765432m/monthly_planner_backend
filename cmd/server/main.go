package main

import (
	"os"
	"os/exec"

	"github.com/098765432m/monthly_planner_backend/internal/app"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	// Init Global logger
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))

	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		zap.S().Fatal("Failed to load .env")
	}
	viper.AutomaticEnv()

	dbUrl := viper.GetString("DB_URL")
	if dbUrl == "" {
		zap.S().Fatal("DB_URL is not set")
	}

	// Call migrate command
	cmd := exec.Command("migrate", "-path", "migrations", "-database", dbUrl, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		zap.S().Fatalf("Migration failed: %v", err)
	}

	if err := app.Run(); err != nil {
		zap.S().Fatal("Applcation failed to run")
	}
}
