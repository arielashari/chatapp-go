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

	// create enum if not exists
	err := DB.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'friend_request_status') THEN
                CREATE TYPE friend_request_status AS ENUM ('pending', 'accepted', 'rejected');
            END IF;
        END$$;
    `).Error
	if err != nil {
		panic("[" + from + "] failed to create enum: " + err.Error())
	}

	// migrate schema
	err = DB.AutoMigrate(dst...)
	if err != nil {
		panic("[" + from + "] failed to migrate database: " + err.Error())
	}

	fmt.Println("[" + from + "] database migrated")
}
