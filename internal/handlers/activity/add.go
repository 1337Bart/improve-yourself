package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

type Handler struct {
	activityService    service.Activity
	dailyReportService service.DailyCheckin
}

func NewHandler(a service.Activity, d service.DailyCheckin) *Handler {
	return &Handler{
		activityService:    a,
		dailyReportService: d,
	}
}

type ActivityLogForm struct {
	Activity  string `form:"activity"`   // Matches the activity name input
	StartTime string `form:"start_time"` // Matches the start time input
	EndTime   string `form:"end_time"`   // Matches the end time input
	Comments  string `form:"comments"`   // Matches the comments input
	Date      string `form:"date"`       // Matches the date input
}

func toServiceActivityLog(input ActivityLogForm) (service.ActivityLog, error) {
	activityLog := service.ActivityLog{}

	startTimeStr := fmt.Sprintf("%sT%s:00", input.Date, input.StartTime)
	startTime, err := parseTime(startTimeStr)
	if err != nil {
		return activityLog, fmt.Errorf("Invalid start time format: %s", err)
	}

	endTimeStr := fmt.Sprintf("%sT%s:00", input.Date, input.EndTime)
	endTime, err := parseTime(endTimeStr)
	if err != nil {
		return activityLog, fmt.Errorf("Invalid end time format: %s", err)
	}

	activityLog.Activity = input.Activity
	activityLog.StartTime = startTime
	activityLog.EndTime = endTime
	activityLog.Comments = input.Comments

	fmt.Println(activityLog)
	return activityLog, nil
}

func parseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	return time.Parse(layout, timeStr)
}

func (h Handler) LogActivityGet(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	todayStr := today.Format("2006-01-02")
	yesterdayStr := yesterday.Format("2006-01-02")

	return render.Render(ctx, views.ActivityLog(todayStr, yesterdayStr))
}

func (h Handler) LogActivityPost(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}
	log.Println(userID, ok)
	input := ActivityLogForm{}
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendString(fmt.Sprintf("<h2>Error parsing input: %v</h2>", err))

	}

	fmt.Printf("input: %+v\n", input)

	serviceACtivityLog, err := toServiceActivityLog(input)
	if err != nil {
		return ctx.SendString("<h2>Error processing your request</h2>")
	}

	err = h.activityService.AddActivityLog(userID, serviceACtivityLog)
	if err != nil {
		return ctx.SendString("<h2>Error processing your request</h2>")
	}

	return ctx.SendString("<p>Activity logged successfully!</p>")
}

type DailyReportForm struct {
	Date               string `form:"date"`
	DidMeditate        string `form:"did_meditate"`
	MinutesOfSports    string `form:"minutes_of_sports"`
	MealsEaten         string `form:"meals_eaten"`
	WaterDrankLiters   string `form:"water_drank_liters"`
	StepsMade          string `form:"steps_made"`
	SleepScore         string `form:"sleep_score"`
	HappinessRating    string `form:"happiness_rating"`
	ProductivityScore  string `form:"productivity_score"`
	StressLevel        string `form:"stress_level"`
	SocialInteractions string `form:"social_interactions"`
	ScreenTimeHours    string `form:"screen_time_hours"`
	WorkHours          string `form:"work_hours"`
	LeisureTimeHours   string `form:"leisure_time_hours"`
	AlcoholUnits       string `form:"alcohol_units"`
	CaffeineCups       string `form:"caffeine_cups"`
	OutdoorTimeHours   string `form:"outdoor_time_hours"`
}

func (h Handler) LogDailyReportPost(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}
	log.Println(userID, ok)
	input := DailyReportForm{}
	if err := ctx.BodyParser(&input); err != nil {
		fmt.Println("err: ", err)
		return ctx.SendString(fmt.Sprintf("<h2>Error parsing input: %v</h2>", err))
	}

	fmt.Printf("input: %+v\n", input)

	serviceDailyReport, err := toServiceDailyReport(input)
	if err != nil {
		return ctx.SendString("<h2>Error converting your input</h2>")
	}

	err = h.dailyReportService.AddDailyReport(userID, serviceDailyReport)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.SendString("<h2>Error processing your request</h2>")
	}

	return ctx.SendString("<p>Daily report logged successfully!</p>")
}

func toServiceDailyReport(input DailyReportForm) (service.ServiceDailyReport, error) {
	var err error
	converted := service.ServiceDailyReport{}

	startTimeStr := fmt.Sprintf("%sT00:01:00", input.Date)

	converted.Date = startTimeStr
	converted.DidMeditate = (input.DidMeditate == "true" || input.DidMeditate == "on")

	if converted.MinutesOfSports, err = strconv.Atoi(input.MinutesOfSports); err != nil {
		return converted, fmt.Errorf("invalid number for minutes of sports: %v", input.MinutesOfSports)
	}
	if converted.MealsEaten, err = strconv.Atoi(input.MealsEaten); err != nil {
		return converted, fmt.Errorf("invalid number for meals eaten: %v", input.MealsEaten)
	}
	if converted.WaterDrankLiters, err = strconv.ParseFloat(input.WaterDrankLiters, 64); err != nil {
		return converted, fmt.Errorf("invalid float for water drank liters: %v", input.WaterDrankLiters)
	}
	if converted.StepsMade, err = strconv.Atoi(input.StepsMade); err != nil {
		return converted, fmt.Errorf("invalid number for steps made: %v", input.StepsMade)
	}
	if converted.SleepScore, err = strconv.Atoi(input.SleepScore); err != nil {
		return converted, fmt.Errorf("invalid number for sleep score: %v", input.SleepScore)
	}
	if converted.HappinessRating, err = strconv.Atoi(input.HappinessRating); err != nil {
		return converted, fmt.Errorf("invalid number for happiness rating: %v", input.HappinessRating)
	}
	if converted.ProductivityScore, err = strconv.Atoi(input.ProductivityScore); err != nil {
		return converted, fmt.Errorf("invalid number for productivity score: %v", input.ProductivityScore)
	}
	if converted.StressLevel, err = strconv.Atoi(input.StressLevel); err != nil {
		return converted, fmt.Errorf("invalid number for stress level: %v", input.StressLevel)
	}
	if converted.SocialInteractions, err = strconv.ParseFloat(input.SocialInteractions, 64); err != nil {
		return converted, fmt.Errorf("invalid float for social interactions: %v", input.SocialInteractions)
	}
	if converted.ScreenTimeHours, err = strconv.ParseFloat(input.ScreenTimeHours, 64); err != nil {
		return converted, fmt.Errorf("invalid float for screen time hours: %v", input.ScreenTimeHours)
	}
	if converted.WorkHours, err = strconv.ParseFloat(input.WorkHours, 64); err != nil {
		return converted, fmt.Errorf("invalid float for work hours: %v", input.WorkHours)
	}
	if converted.LeisureTimeHours, err = strconv.ParseFloat(input.LeisureTimeHours, 64); err != nil {
		return converted, fmt.Errorf("invalid float for leisure time hours: %v", input.LeisureTimeHours)
	}
	if converted.AlcoholUnits, err = strconv.ParseFloat(input.AlcoholUnits, 64); err != nil {
		return converted, fmt.Errorf("invalid float for alcohol units: %v", input.AlcoholUnits)
	}
	if converted.CaffeineCups, err = strconv.ParseFloat(input.CaffeineCups, 64); err != nil {
		return converted, fmt.Errorf("invalid float for caffeine cups: %v", input.CaffeineCups)
	}
	if converted.OutdoorTimeHours, err = strconv.ParseFloat(input.OutdoorTimeHours, 64); err != nil {
		return converted, fmt.Errorf("invalid float for outdoor time hours: %v", input.OutdoorTimeHours)
	}

	return converted, nil
}
func (h Handler) LogDailyReportGet(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	todayStr := today.Format("2006-01-02")
	yesterdayStr := yesterday.Format("2006-01-02")

	return render.Render(ctx, views.DailyReportLog(todayStr, yesterdayStr))
}
