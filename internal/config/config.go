package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	ParseTime string `mapstructure:"parse_time"`
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

var AppGlobalConfigData Config

func InitConfig() error {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") //Env file in Root

	if err := viper.ReadInConfig(); err != nil {
		zap.S().Fatal("Cannot read config file !")
	}

	if err := viper.Unmarshal(&AppGlobalConfigData); err != nil {
		zap.S().Fatal("Cannot unmarshal from config file !")
	}

	return nil
}
