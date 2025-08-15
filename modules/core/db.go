package core

import (
	"fmt"
	"go-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := config.App.DatabaseDSN
	connection := postgres.Open(dsn)

	logLevel := logger.Silent

	if config.App.ApplicationEnvironment == "dev" {
		logLevel = logger.Info
	}

	database, err := gorm.Open(connection, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic("[database] failed to connect database")
	}

	DB = database
}

func AutoMigrate(from string, dst ...interface{}) {
	if !config.App.DbMigrate {
		return
	}

	err := DB.AutoMigrate(dst...)

	if err != nil {
		panic("[" + from + "] failed to migrate database")
	}

	fmt.Println("[" + from + "] database migrated")
}
