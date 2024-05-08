package service

import "time"

type ActivityLog struct {
	Activity  string
	StartTime time.Time
	EndTime   time.Time
	Comments  string
}

type Activity interface {
	AddActivityLog(id string, activityLog ActivityLog) error
}
