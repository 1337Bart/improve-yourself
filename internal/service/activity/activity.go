package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/service"
	"gorm.io/gorm"
	"math"
	"sort"
	"strconv"
	"time"
)

type Activity struct {
	SqlDb *gorm.DB
}

func NewActivityService(sqlDbConn *gorm.DB) *Activity {
	return &Activity{
		SqlDb: sqlDbConn,
	}
}

func (a *Activity) AddActivityLog(userID string, activity service.ActivityLog) error {
	tx := a.SqlDb.Model(&model.ActivityLog{}).Create(map[string]interface{}{
		"uuid":       userID,
		"activity":   activity.Activity,
		"start_time": activity.StartTime,
		"end_time":   activity.EndTime,
		"comments":   activity.Comments,
	})

	err := tx.Error
	if err != nil {
		return fmt.Errorf("error updating record: %v", err)
	}

	return err

}

func (a *Activity) GetActivitiesForDay(userID, date string) ([]service.ActivityLogDisplay, error) {
	var activities []service.ActivityLog

	startOfDay, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)

	result := a.SqlDb.Model(&model.ActivityLog{}).
		Where("uuid = ? AND start_time >= ? AND end_time <= ?", userID, startOfDay, endOfDay).
		Find(&activities)

	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving activities: %v", result.Error)
	}

	return toDisplayctivities(activities), nil
}

// todo - pozbyć się tego (robie to samo znowu w handlerze)
func toDisplayctivities(activities []service.ActivityLog) []service.ActivityLogDisplay {
	displayActivities := make([]service.ActivityLogDisplay, 0, len(activities))

	for _, item := range activities {
		displayActivities = append(displayActivities, service.ActivityLogDisplay{
			Activity:  item.Activity,
			StartTime: item.StartTime.Format("Jan 2, 2006 15:04"),
			EndTime:   item.EndTime.Format("Jan 2, 2006 15:04"),
			Duration:  fmt.Sprintf("%v", item.EndTime.Sub(item.StartTime).Minutes()),
			Comments:  item.Comments,
		})
	}

	return displayActivities
}

func (a *Activity) GetActivityDistributionByPeriod(userId, startDate, endDate string) (map[string]int, error) {
	var activities []service.ActivityLog

	startDay, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	endDay, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	// add full one day to the end date to include the end date in the query
	endDay = endDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := a.SqlDb.Model(&model.ActivityLog{}).
		Where("uuid = ? AND start_time >= ? AND end_time <= ?", userId, startDay, endDay).
		Find(&activities)

	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving activities: %v", result.Error)
	}

	dayCount := make(map[string]int)
	for _, activity := range activities {
		day := activity.StartTime.Weekday().String()
		dayCount[day]++
	}

	return dayCount, nil
}

func (a *Activity) GetLongestActivityByPeriod(userId, startDate, endDate string) (service.ActivityLogDisplay, error) {
	var longestActivity service.ActivityLogDisplay
	var activities []service.ActivityLog

	startDay, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return longestActivity, fmt.Errorf("error parsing date: %v", err)
	}

	endDay, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return longestActivity, fmt.Errorf("error parsing date: %v", err)
	}

	// add full one day to the end date to include the end date in the query
	endDay = endDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := a.SqlDb.Model(&model.ActivityLog{}).
		Where("uuid = ? AND start_time >= ? AND end_time <= ?", userId, startDay, endDay).
		Find(&activities)

	if result.Error != nil {
		return longestActivity, fmt.Errorf("error retrieving activities: %v", result.Error)
	}

	maxDuration := time.Duration(0)
	for _, activity := range activities {
		duration := activity.EndTime.Sub(activity.StartTime)
		if duration > maxDuration {
			maxDuration = duration
			longestActivity = toDisplayctivity(activity)
		}
	}
	return longestActivity, nil
}

func toDisplayctivity(activity service.ActivityLog) service.ActivityLogDisplay {
	return service.ActivityLogDisplay{
		Activity:  activity.Activity,
		StartTime: activity.StartTime.Format("Jan 2, 2006 15:04"),
		EndTime:   activity.EndTime.Format("Jan 2, 2006 15:04"),
		Duration:  fmt.Sprintf("%v", math.Abs(activity.EndTime.Sub(activity.StartTime).Minutes())),
		Comments:  activity.Comments,
	}
}

func (a *Activity) GetTopThreeActivities(userId, startDate, endDate string) (map[string]int, error) {
	var activities []service.ActivityLog

	startDay, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	endDay, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	// add full one day to the end date to include the end date in the query
	endDay = endDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := a.SqlDb.Model(&model.ActivityLog{}).
		Where("uuid = ? AND start_time >= ? AND end_time <= ?", userId, startDay, endDay).
		Find(&activities)

	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving activities: %v", result.Error)
	}

	convertedAcitivites := toDisplayctivities(activities)

	activityToDuration := make(map[string]int)
	for _, activity := range convertedAcitivites {
		duration, err := strconv.Atoi(activity.Duration)
		if err != nil {
			return nil, fmt.Errorf("error converting duration to int %v", err)
		}

		activityToDuration[activity.Activity] += duration
	}

	return getTopThreeActivities(activityToDuration), nil
}

func getTopThreeActivities(activityToDuration map[string]int) map[string]int {
	type kv struct {
		Key   string
		Value int
	}

	var kvPairs []kv
	for k, v := range activityToDuration {
		kvPairs = append(kvPairs, kv{k, v})
	}

	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i].Value > kvPairs[j].Value
	})

	topThree := make(map[string]int)
	for i := 0; i < 3 && i < len(kvPairs); i++ {
		topThree[kvPairs[i].Key] = kvPairs[i].Value
	}

	return topThree
}

func (a *Activity) GetTimeTrackedAveragesDaily(userId, startDate, endDate string) ([]service.DayCoverage, error) {
	var activities []model.ActivityLog
	trackedPercentPerDay := make(map[string]int)

	startDay, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	endDay, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	// Add full one day to the end date to include the end date in the query
	endDay = endDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := a.SqlDb.Model(&model.ActivityLog{}).
		Where("uuid = ? AND start_time >= ? AND end_time <= ?", userId, startDay, endDay).
		Find(&activities)

	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving activities: %v", result.Error)
	}

	for d := startDay; d.Before(endDay); d = d.AddDate(0, 0, 1) {
		trackedPercentPerDay[d.Format("2006-01-02")] = 0
	}

	for _, activity := range activities {
		start := activity.StartTime
		end := activity.EndTime

		for start.Before(end) {
			currentDay := start.Format("2006-01-02")
			nextDay := start.AddDate(0, 0, 1).Truncate(24 * time.Hour)

			var duration time.Duration
			if end.Before(nextDay) {
				duration = end.Sub(start)
			} else {
				duration = nextDay.Sub(start)
			}

			trackedHours := duration.Hours()
			trackedPercentPerDay[currentDay] += int(math.Round(trackedHours / 24.0 * 100))

			start = nextDay
		}
	}

	var totalTrackedPercent int
	var dayCount int

	// Convert map to slice of service.DayCoverage
	var dayCoverages []service.DayCoverage
	for day, percent := range trackedPercentPerDay {
		dayCoverages = append(dayCoverages, service.DayCoverage{Day: day, Coverage: percent})
		totalTrackedPercent += percent
		dayCount++
	}

	if dayCount > 0 {
		dayCoverages = append(dayCoverages, service.DayCoverage{Day: "weekly average", Coverage: int(math.Round(float64(totalTrackedPercent) / float64(dayCount)))})
	} else {
		dayCoverages = append(dayCoverages, service.DayCoverage{Day: "weekly average", Coverage: 0})
	}

	sort.Slice(dayCoverages, func(i, j int) bool {
		return dayCoverages[i].Day > dayCoverages[j].Day
	})

	return dayCoverages, nil
}
