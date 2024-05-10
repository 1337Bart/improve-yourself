package model

import "time"

type ActivityLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserUUID  string    `gorm:"type:uuid;column:uuid" json:"uuid"`
	Activity  string    `gorm:"column:activity" json:"activity"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Comments  string    `gorm:"column:comments;type:text" json:"comments"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
