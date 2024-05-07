package model

import "time"

type Settings struct {
	ID        uint      `gorm:"autoIncrement;column:id" json:"id"`
	UUID      string    `gorm:"primaryKey;type:uuid;uniqueIndex;column:uuid" json:"uuid"`
	SearchOn  bool      `json:"searchOn"`
	AddNew    bool      `json:"addNew"`
	Amount    uint      `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}
