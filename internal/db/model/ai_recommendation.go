package model

type AiRecommendation struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserUUID           string `gorm:"type:uuid;column:uuid" json:"uuid"`
	Insights           string `gorm:"type:text;column:insights"`
	ProposedActivities string `gorm:"type:text;column:proposed_activities"`
}
