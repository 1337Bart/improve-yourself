package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"time"
)

func groupActivitiesByDate(activities []service.ActivityLogDisplay) map[string][]service.ActivityLogDisplayTransformed {
	groupedActivities := make(map[string][]service.ActivityLogDisplayTransformed)

	dateFormats := []string{
		"January 02, 2006 15:04",
		"Jan 2, 2006 15:04",
		"2006-01-02 15:04",
		"2006/01/02 15:04",
		"02.01.2006 15:04",
	}

	for _, activity := range activities {
		var startTime, endTime time.Time
		var err error

		for _, format := range dateFormats {
			startTime, err = time.Parse(format, activity.StartTime)
			if err == nil {
				break
			}
		}
		if err != nil {
			continue
		}

		for _, format := range dateFormats {
			endTime, err = time.Parse(format, activity.EndTime)
			if err == nil {
				break
			}
		}
		if err != nil {
			continue
		}

		date := startTime.Format("02.01.2006")

		startTimeStr := startTime.Format("15:04")
		endTimeStr := endTime.Format("15:04")

		transformedActivity := service.ActivityLogDisplayTransformed{
			Activity:  activity.Activity,
			StartTime: startTimeStr,
			EndTime:   endTimeStr,
			Duration:  activity.Duration,
			Comments:  activity.Comments,
		}

		groupedActivities[date] = append(groupedActivities[date], transformedActivity)
	}

	fmt.Printf("groupActivitiesByDate: %+v", groupedActivities)
	return groupedActivities
}

func (h *Handler) ActivitiesForDayGet(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}

	date := ctx.Query("selected_date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	activities, err := h.activityService.GetActivitiesForDay(userID, date)
	if err != nil {
		return ctx.Status(500).SendString(fmt.Sprintf("<h2>Error fetching activities: %v</h2>", err))
	}
	fmt.Printf("ActivitiesForDayGet activities: %+v\n", activities)
	groupedActivities := groupActivitiesByDate(activities)

	return render.Render(ctx, views.ActivityDayLog(groupedActivities))
}

func (h *Handler) ActivitiesForDayPost(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}

	date := ctx.FormValue("selected_date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	activities, err := h.activityService.GetActivitiesForDay(userID, date)
	if err != nil {
		return ctx.Status(500).SendString(fmt.Sprintf("<h2>Error fetching activities: %v</h2>", err))
	}

	fmt.Printf("ActivitiesForDayPost activities: %+v\n", activities)
	groupedActivities := groupActivitiesByDate(activities)
	return render.Render(ctx, views.ActivityDayLog(groupedActivities))
}

func (h *Handler) DailyCheckinForDayGet(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}

	date := ctx.Query("selected_date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	checkin, _ := h.dailyReportService.GetDailyCheckinForDay(userID, date)

	return render.Render(ctx, views.DailyCheckinGet(ToDailyReportForm(checkin)))
}

func ToDailyReportForm(serviceDailyReport service.ServiceDailyReport) service.DailyReportForm {
	ret := service.DailyReportForm{}

	if serviceDailyReport.Date != "" {
		ret.Date = serviceDailyReport.Date[:10]
	} else {
		ret.Date = "Entries not found for selected date"
	}

	ret.DidMeditate = fmt.Sprintf("%t", serviceDailyReport.DidMeditate)
	ret.MinutesOfSports = fmt.Sprintf("%d", serviceDailyReport.MinutesOfSports)
	ret.MealsEaten = fmt.Sprintf("%d", serviceDailyReport.MealsEaten)
	ret.WaterDrankLiters = fmt.Sprintf("%.2f", serviceDailyReport.WaterDrankLiters)
	ret.StepsMade = fmt.Sprintf("%d", serviceDailyReport.StepsMade)
	ret.SleepScore = fmt.Sprintf("%d", serviceDailyReport.SleepScore)
	ret.HappinessRating = fmt.Sprintf("%d", serviceDailyReport.HappinessRating)
	ret.ProductivityScore = fmt.Sprintf("%d", serviceDailyReport.ProductivityScore)
	ret.StressLevel = fmt.Sprintf("%d", serviceDailyReport.StressLevel)
	ret.SocialInteractions = fmt.Sprintf("%.2f", serviceDailyReport.SocialInteractions)
	ret.ScreenTimeHours = fmt.Sprintf("%.2f", serviceDailyReport.ScreenTimeHours)
	ret.WorkHours = fmt.Sprintf("%.2f", serviceDailyReport.WorkHours)
	ret.LeisureTimeHours = fmt.Sprintf("%.2f", serviceDailyReport.LeisureTimeHours)
	ret.AlcoholUnits = fmt.Sprintf("%.2f", serviceDailyReport.AlcoholUnits)
	ret.CaffeineCups = fmt.Sprintf("%.2f", serviceDailyReport.CaffeineCups)
	ret.OutdoorTimeHours = fmt.Sprintf("%.2f", serviceDailyReport.OutdoorTimeHours)

	return ret
}
