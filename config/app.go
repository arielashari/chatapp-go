package config

import (
	"github.com/spf13/viper"
	"log"
)

type Environment struct {
	ApplicationPort        string `mapstructure:"APPLICATION_PORT"`
	DatabaseDSN            string `mapstructure:"DATABASE_DSN"`
	ApplicationEnvironment string `mapstructure:"APP_ENV"`
	JwtSecret              string `mapstructure:"JWT_SECRET"`
	DbMigrate              bool   `mapstructure:"DB_MIGRATE"`
}

var App Environment

func LoadConfig() {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	setDefault()

	err := viper.ReadInConfig()

	err = viper.Unmarshal(&App)

	if err != nil {
		log.Fatal(err)
	}

	validateConfig()
}

func setDefault() {
	viper.SetDefault("APPLICATION_PORT", "")
	viper.SetDefault("DATABASE_DSN", "")
	viper.SetDefault("DB_MIGRATE", false)
	viper.SetDefault("APP_ENV", "dev")
}

func validateConfig() {
	if App.ApplicationPort == "" {
		log.Fatalf("No application port supplied")
	}
	if App.DatabaseDSN == "" {
		log.Fatalf("No database DSN supplied")
	}
	if App.JwtSecret == "" {
		log.Fatalf("No JWT secret supplied")
	}
}
