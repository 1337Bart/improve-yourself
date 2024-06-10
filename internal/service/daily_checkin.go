package service

type ServiceDailyReport struct {
	Date               string
	DidMeditate        bool
	MinutesOfSports    int
	MealsEaten         int
	WaterDrankLiters   float64
	StepsMade          int
	SleepScore         int
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
	AddDailyCheckin(string, ServiceDailyReport) error
	GetDailyCheckinForDay(string, string) (ServiceDailyReport, error)
}

type DailyReportForm struct {
	Date               string
	DidMeditate        string
	MinutesOfSports    string
	MealsEaten         string
	WaterDrankLiters   string
	StepsMade          string
	SleepScore         string
	HappinessRating    string
	ProductivityScore  string
	StressLevel        string
	SocialInteractions string
	ScreenTimeHours    string
	WorkHours          string
	LeisureTimeHours   string
	AlcoholUnits       string
	CaffeineCups       string
	OutdoorTimeHours   string
}
