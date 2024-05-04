package db

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) (*gorm.DB, error) {
	if dbUrl == "" {
		return nil, fmt.Errorf("empty db url env")
	}

	DbConn, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}

	err = DbConn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		return nil, fmt.Errorf("unable to create uuid extension: %w", err)
	}

	err = DbConn.AutoMigrate(&model.User{}, &model.Settings{}, &model.PotatoTime{})
	if err != nil {
		return nil, fmt.Errorf("unable to migrate DB: %w", err)
	}

	err = DbConn.Exec("ALTER TABLE settings ADD COLUMN IF NOT EXISTS id UUID DEFAULT uuid_generate_v4();").Error
	if err != nil {
		return nil, fmt.Errorf("unable to alter table to add new id column: %w", err)
	}

	return DbConn, nil
}
