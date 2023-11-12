package config

import (
	"github.com/Lyalyashechka/actresses/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Logger logger.Config
}

type Server struct {
	Host string
	Port string
}

func LoadConfig(config *Config, path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}
