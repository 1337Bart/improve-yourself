package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// todo: this should not be a global var
var DbConn *gorm.DB
var err error

func InitDB(dbUrl string) error {
	if dbUrl == "" {
		return fmt.Errorf("empty db url env")
	}

	DbConn, err = gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		return fmt.Errorf("unable to connect to DB: %w", err)
	}

	err = DbConn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		return fmt.Errorf("unable to create uuid extension: %w: ", err)
	}

	err = DbConn.AutoMigrate(&User{}, &Settings{})
	if err != nil {
		return fmt.Errorf("unable to migrate DB: %w", err)
	}

	return nil
}
