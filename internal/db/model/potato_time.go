package model

import "time"

type PotatoTime struct {
	ID           uint      `gorm:"autoIncrement;column:id" json:"id"`
	UUID         string    `gorm:"primaryKey;type:uuid;uniqueIndex;column:uuid" json:"uuid"`
	Amount       int       `json:"amount"`
	UpdatesCount uint      `json:"updatesCount"`
	TotalAdded   uint      `json:"totalAdded"`
	TotalUsed    uint      `json:"totalUsed"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
