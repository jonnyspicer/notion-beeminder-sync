package utils

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
