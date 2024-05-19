package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/service"
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	SqlDb *gorm.DB
}

func NewActviyService(sqlDbConn *gorm.DB) *Activity {
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
