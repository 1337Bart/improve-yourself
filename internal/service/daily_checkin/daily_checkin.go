package daily_checkin

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/service"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type DailyCheckin struct {
	SqlDb *gorm.DB
}

func NewDailyCheckinService(sqlDbConn *gorm.DB) *DailyCheckin {
	return &DailyCheckin{
		SqlDb: sqlDbConn,
	}
}

func (d *DailyCheckin) AddDailyCheckin(userID string, checkin service.ServiceDailyReport) error {
	date, err := parseTime(checkin.Date)
	if err != nil {
		return fmt.Errorf("invalid start time format: %s", err)
	}

	tx := d.SqlDb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}, {Name: "date"}},
		DoUpdates: clause.AssignmentColumns([]string{"date", "did_meditate", "minutes_of_sports", "meals_eaten", "water_drank_liters", "steps_made", "sleep_score", "happiness_rating", "productivity_score", "stress_level", "social_interactions", "screen_time_hours", "work_hours", "leisure_time_hours", "alcohol_units", "caffeine_cups", "outdoor_time_hours"}),
	}).Create(&model.DailyCheckIn{
		UserUUID:           userID,
		Date:               date,
		DidMeditate:        checkin.DidMeditate,
		MinutesOfSports:    checkin.MinutesOfSports,
		MealsEaten:         checkin.MealsEaten,
		WaterDrankLiters:   checkin.WaterDrankLiters,
		StepsMade:          checkin.StepsMade,
		SleepScore:         checkin.SleepScore,
		HappinessRating:    checkin.HappinessRating,
		ProductivityScore:  checkin.ProductivityScore,
		StressLevel:        checkin.StressLevel,
		SocialInteractions: checkin.SocialInteractions,
		ScreenTimeHours:    checkin.ScreenTimeHours,
		WorkHours:          checkin.WorkHours,
		LeisureTimeHours:   checkin.LeisureTimeHours,
		AlcoholUnits:       checkin.AlcoholUnits,
		CaffeineCups:       checkin.CaffeineCups,
		OutdoorTimeHours:   checkin.OutdoorTimeHours,
	})
	err = tx.Error
	if err != nil {
		return fmt.Errorf("error upserting record: %v", err)
	}

	return err
}

func parseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	return time.Parse(layout, timeStr)
}

func (d *DailyCheckin) GetDailyCheckinForDay(userID string, date string) (service.ServiceDailyReport, error) {
	var checkin model.DailyCheckIn

	datePattern := date + "%"

	err := d.SqlDb.Where("uuid = ? AND to_char(date, 'YYYY-MM-DD') LIKE ?", userID, datePattern).First(&checkin).Error
	if err != nil {
		return service.ServiceDailyReport{}, fmt.Errorf("error fetching daily checkin: %v", err)
	}

	fmt.Printf("checkin: %+v\n", checkin)

	return service.ServiceDailyReport{
		Date:               checkin.Date.Format("2006-01-02T15:04:05"),
		DidMeditate:        checkin.DidMeditate,
		MinutesOfSports:    checkin.MinutesOfSports,
		MealsEaten:         checkin.MealsEaten,
		WaterDrankLiters:   checkin.WaterDrankLiters,
		StepsMade:          checkin.StepsMade,
		SleepScore:         checkin.SleepScore,
		HappinessRating:    checkin.HappinessRating,
		ProductivityScore:  checkin.ProductivityScore,
		StressLevel:        checkin.StressLevel,
		SocialInteractions: checkin.SocialInteractions,
		ScreenTimeHours:    checkin.ScreenTimeHours,
		WorkHours:          checkin.WorkHours,
		LeisureTimeHours:   checkin.LeisureTimeHours,
		AlcoholUnits:       checkin.AlcoholUnits,
		CaffeineCups:       checkin.CaffeineCups,
		OutdoorTimeHours:   checkin.OutdoorTimeHours,
	}, nil
}
