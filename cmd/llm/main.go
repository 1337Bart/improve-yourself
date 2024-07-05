package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/internal/service/activity"
	"github.com/1337Bart/improve-yourself/internal/service/ai_insight"
	"github.com/1337Bart/improve-yourself/internal/service/daily_checkin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
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
	reports := fetchReports(activityService, dailyReportService, userID, dates)

	fileName := "reports.json"
	err = saveReportsToFile(reports, fileName)
	if err != nil {
		log.Fatalf("error saving reports to file: %s", err)
	}

	output, err := executePythonScript(fileName)
	if err != nil {
		log.Fatalf("error executing Python script: %s", err)
	}

	var scriptOutput service.AiRecommendation

	err = json.Unmarshal([]byte(output), &scriptOutput)
	if err != nil {
		log.Fatalf("error unmarshalling JSON output: %s", err)
	}

	aiService := ai_insight.NewAiInsightService(dbConn)

	err = aiService.AddAiInsight(userID, scriptOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("AI insights saved successfully. Exiting..")
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

func fetchReports(activityService *activity.Activity, dailyReportService *daily_checkin.DailyCheckin, userID string, dates []string) dailyReportMap {
	reports := make(dailyReportMap)
	for _, date := range dates {

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
	return reports
}

func saveReportsToFile(reports dailyReportMap, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(reports)
}

func executePythonScript(fileName string) (string, error) {
	cmd := exec.Command("python3", "llm_script.py", fileName)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	fmt.Println("out.String: ", out.String())
	return out.String(), nil
}
