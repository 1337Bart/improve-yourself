package model

import "time"

type ActivityLog struct {
	ID        uint      `gorm:"autoIncrement;column:id" json:"id"` // Unique ID for the activity record
	UUID      string    `gorm:"primaryKey;type:uuid;uniqueIndex;column:uuid" json:"uuid"`
	Activity  string    `gorm:"column:activity" json:"activity"`           // Description of the activity (e.g., "breakfast", "morning run")
	Duration  uint      `gorm:"column:duration" json:"duration"`           // Duration of the activity in minutes
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`       // Start time of the activity
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`           // End time of the activity
	Comments  string    `gorm:"column:comments;type:text" json:"comments"` // Optional comments about the activity
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`       // Timestamp of when the record was created
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`       // Timestamp of the last update to the record
}
