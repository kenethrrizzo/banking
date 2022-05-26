package config

import (
	"fmt"

	"github.com/kenethrrizzo/banking/logger"
	"github.com/spf13/viper"
)

func InitializeDatabaseConfig() {
	viper.SetConfigName("database-config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.Debug("Database config loaded")
}