package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var State = "DEV"
var RenewDB = true
var prodFile = "prodcat_prod"
var devFile = "prodcat_dev"

func InitConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./config")
	config.AddConfigPath("$HOME")
	config.SetEnvPrefix("prodcat")
	viper.BindEnv("state")

	// $env:PRODCAT_STATE = 'PRODUCTION'
	currentEnv := os.Getenv("PRODCAT_STATE")

	if currentEnv == "" || currentEnv == "DEV" {
		State = "DEV"
		log.Println("--- DEV CONFIG LOADED ---")
		config.SetConfigName(devFile)
	}

	if currentEnv == "PROD" || currentEnv == "PRODUCTION" {
		State = "PRODUCTION"
		log.Println("---! PRODUCTION CONFIG LOADED !---")
		config.SetConfigName(prodFile)
	}

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}
	return config
}
