package model

import "time"

type Settings struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	SearchOn  bool      `json:"searchOn"`
	AddNew    bool      `json:"addNew"`
	Amount    uint      `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}
