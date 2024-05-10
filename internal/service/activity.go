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

type Activity interface {
	AddActivityLog(id string, activityLog ActivityLog) error
	GetActivitiesForDay(userID, date string) ([]ActivityLogDisplay, error)
}
