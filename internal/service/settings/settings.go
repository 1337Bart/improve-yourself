package settings

import (
	"fmt"
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
	fmt.Println("entered get")
	err = s.SqlDb.Where("id = ?", id).First(&settings).Error
	fmt.Println("err: ", err)
	return settings, err
}

func (s *Settings) Update(settings *model.Settings) error {
	tx := s.SqlDb.Select("search_on", "add_new", "amount", "updated_at").Where("id = ?", settings.ID).Updates(settings)
	return tx.Error
}
