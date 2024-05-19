package service

import "time"

type ActivityLog struct {
	Activity  string
	StartTime time.Time
	EndTime   time.Time
	Comments  string
}

type ActivityLogDisplay struct {
	Activity  string
	StartTime string
	EndTime   string
	Duration  string
	Comments  string
}

type ActivityLogDisplayTransformed struct {
	Activity  string
	StartTime string
	EndTime   string
	Duration  string
	Comments  string
}

type DayCoverage struct {
	Day      string
	Coverage int
}

type Activity interface {
	AddActivityLog(id string, activityLog ActivityLog) error
	GetActivitiesForDay(userID, date string) ([]ActivityLogDisplay, error)

	GetActivityDistributionByPeriod(userId, startDate, endDate string) (map[string]int, error)
	GetLongestActivityByPeriod(userId, startDate, endDate string) (ActivityLogDisplay, error)
	GetTopThreeActivities(userId, startDate, endDate string) (map[string]int, error)
	GetTimeTrackedAveragesDaily(userId, startDate, endDate string) ([]DayCoverage, error)
}
