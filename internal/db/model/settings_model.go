package model

import "time"

type Settings struct {
	ID        string    `gorm:"type:uuid;primarykey;column:id" json:"id"`
	SearchOn  bool      `json:"searchOn"`
	AddNew    bool      `json:"addNew"`
	Amount    uint      `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
}
