package config

import (
	"fmt"

	"github.com/kenethrrizzo/banking/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Domain   string `mapstructure:"domain"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func NewDatabaseConfig() DatabaseConfig {
	return unmarshalConfig().Database
}

func unmarshalConfig() Config {
	var config Config
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath("./config")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.Debug("Database config loaded")

	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode into struct, %v", err))
	}
	logger.Debug("Database config unmarshalled")

	return config
}