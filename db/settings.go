package db

import "time"

type Settings struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	SearchOn  bool      `json:"searchOn"`
	AddNew    bool      `json:"addNew"`
	Amount    uint      `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Settings) Get() error {
	err := DbConn.Where("id = 1").First(s).Error
	return err
}

func (s *Settings) Update() error {
	tx := DbConn.Select("search_on", "add_new", "amount", "updated_at").Where("id = 1").Updates(s)
	return tx.Error
}
