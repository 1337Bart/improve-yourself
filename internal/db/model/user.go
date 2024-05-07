package model

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	IsAdmin   bool      `gorm:"default:false" json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
