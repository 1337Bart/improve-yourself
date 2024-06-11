package main

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/internal/service/activity"
	"github.com/1337Bart/improve-yourself/internal/service/daily_checkin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type CombinedDailyReport struct {
	DailyReport service.ServiceDailyReport
	Activities  []service.ActivityLogDisplay
}

type dailyReportMap map[string]CombinedDailyReport

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	dbConn, err := db.InitDB(dbUrl)
	if err != nil {
		fmt.Printf("error connecting to database: %s\n", err)
		panic("error connecting to db")
	}

	activityService := activity.NewActivityService(dbConn)
	dailyReportService := daily_checkin.NewDailyCheckinService(dbConn)

	userID := "afcd065f-7205-4353-b7f7-f51d3f0abff3"

	dates := getLastNDays(11)
	reports := make(dailyReportMap)
	for _, date := range dates {
		fmt.Println("####DATE####", date)
		dailyReport, err := dailyReportService.GetDailyCheckinForDay(userID, date)
		if err != nil {
			log.Printf("error fetching daily report for date %s: %s\n", date, err)
			continue
		}

		activities, err := activityService.GetActivitiesForDay(userID, date)
		if err != nil {
			log.Printf("error fetching activities for date %s: %s\n", date, err)
			continue
		}

		combinedReport := CombinedDailyReport{
			DailyReport: dailyReport,
			Activities:  activities,
		}

		reports[date] = combinedReport
	}
}

func getLastNDays(N int) []string {
	var dates []string
	now := time.Now()
	for i := N - 1; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		dates = append(dates, date)
	}
	return dates
}
