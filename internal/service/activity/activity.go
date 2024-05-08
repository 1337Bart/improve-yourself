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
		"duration":   calculateDuration(activity.StartTime, activity.EndTime),
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

func calculateDuration(startTime time.Time, endTime time.Time) uint {
	return uint(endTime.Sub(startTime).Minutes())
}
