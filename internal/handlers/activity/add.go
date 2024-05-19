package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type Handler struct {
	activityService service.Activity
}

func NewHandler(a service.Activity) *Handler {
	return &Handler{
		activityService: a,
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
	fmt.Println("at activity log")
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
