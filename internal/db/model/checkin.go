package model

import "time"

type DailyCheckIn struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserUUID           string    `gorm:"type:uuid;column:uuid;uniqueIndex:idx_user_date" json:"uuid"`
	Date               time.Time `gorm:"column:date;uniqueIndex:idx_user_date" json:"date"`
	DidMeditate        bool      `gorm:"column:did_meditate" json:"did_meditate"`
	MinutesOfSports    int       `gorm:"column:minutes_of_sports" json:"minutes_of_sports"`
	MealsEaten         int       `json:"meals_eaten"`
	WaterDrankLiters   float64   `json:"water_drank_liters"`
	StepsMade          int       `json:"steps_made"`
	SleepScore         int       `json:"sleep_score"`
	HappinessRating    int       `json:"happiness_rating"`
	ProductivityScore  int       `json:"productivity_score"`
	StressLevel        int       `json:"stress_level"`
	SocialInteractions float64   `json:"social_interactions"`
	ScreenTimeHours    float64   `json:"screen_time_hours"`
	WorkHours          float64   `json:"work_hours"`
	LeisureTimeHours   float64   `json:"leisure_time_hours"`
	AlcoholUnits       float64   `json:"alcohol_units"`
	CaffeineCups       float64   `json:"caffeine_cups"`
	OutdoorTimeHours   float64   `json:"outdoor_time_hours"`
}
