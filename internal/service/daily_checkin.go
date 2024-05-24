package service

type ServiceDailyReport struct {
	Date               string
	DidMeditate        bool
	MinutesOfSports    int
	MealsEaten         int
	WaterDrankLiters   float64
	StepsMade          int
	SleepScore         int
	DominatingEmotion  string
	HappinessRating    int
	ProductivityScore  int
	StressLevel        int
	SocialInteractions float64
	ScreenTimeHours    float64
	WorkHours          float64
	LeisureTimeHours   float64
	AlcoholUnits       float64
	CaffeineCups       float64
	OutdoorTimeHours   float64
}

type DailyCheckin interface {
	AddDailyReport(string, ServiceDailyReport) error
}
