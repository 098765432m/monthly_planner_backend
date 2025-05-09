package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	App AppConfig `mapstructure:"app"`
}

var AppGlobalConfigData Config

func LoadConfig() error {

	// Load .env file
	// if err := godotenv.Load(".env"); err != nil {
	// 	zap.S().Warnln("No .env file found or failed to load it.")
	// }

	// Load YAML config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") //Env file in Root

	if err := viper.ReadInConfig(); err != nil {
		zap.S().Fatal("Cannot read config file !")
	}

	// Enable reading evironment variables
	// viper.AutomaticEnv()

	if err := viper.Unmarshal(&AppGlobalConfigData); err != nil {
		zap.S().Fatal("Cannot unmarshal from config file !")
	}

	return nil
}
