package settings

import (
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"gorm.io/gorm"
)

type Settings struct {
	SqlDb *gorm.DB
}

func NewSettingsService(sqlDbConn *gorm.DB) *Settings {
	return &Settings{
		SqlDb: sqlDbConn,
	}
}

func (s *Settings) Get(id string) (settings *model.Settings, err error) {
	err = s.SqlDb.Where("uuid = ?", id).First(&settings).Error
	return settings, err
}

func (s *Settings) Update(settings *model.Settings) error {
	tx := s.SqlDb.Select("search_on", "add_new", "amount", "updated_at").Where("uuid = ?", settings.UUID).Updates(settings)
	return tx.Error
}

func (s *Settings) CreateDefault(userID string) error {
	settings := &model.Settings{
		UUID:     userID,
		Amount:   uint(10),
		SearchOn: true,
		AddNew:   true,
	}

	tx := s.SqlDb.Create(settings)
	return tx.Error
}
