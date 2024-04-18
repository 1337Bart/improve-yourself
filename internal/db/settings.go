package db

import "github.com/1337Bart/improve-yourself/internal/db/model"

func (s *model.Settings) Get() error {
	err := DbConn.Where("id = 1").First(s).Error
	return err
}

func (s *model.Settings) Update() error {
	tx := DbConn.Select("search_on", "add_new", "amount", "updated_at").Where("id = 1").Updates(s)
	return tx.Error
}
