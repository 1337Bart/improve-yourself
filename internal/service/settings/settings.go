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

// TODO modify this to use actual ID - handler needs to pass ID to this
func (s *Settings) Get(id string, settings *model.Settings) error {
	err := s.SqlDb.Where("id = ?", id).First(settings).Error
	return err
}

func (s *Settings) Update(id string, settings *model.Settings) error {
	tx := s.SqlDb.Select("search_on", "add_new", "amount", "updated_at").Where("id = ?", id).Updates(settings)
	return tx.Error
}
