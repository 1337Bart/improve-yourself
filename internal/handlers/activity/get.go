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

	for _, activity := range activities {
		// Parse the StartTime and EndTime to time.Time
		startTime, err := time.Parse("January 02, 2006 15:04", activity.StartTime)
		if err != nil {
			continue // Skip if there's an error in parsing
		}

		endTime, err := time.Parse("January 02, 2006 15:04", activity.EndTime)
		if err != nil {
			continue // Skip if there's an error in parsing
		}

		// Extract date in "DD:MM:YYYY" format
		date := startTime.Format("02.01.2006")

		// Extract time in "HH:MM" format
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
	fmt.Printf("activities: %+v\n", activities)
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
	} else { // od else do wywalenia caly blok
		selectedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return ctx.Status(400).SendString("<h2>Error: Invalid date format</h2>")
		}
		if time.Since(selectedDate).Hours() > 7*24 {
			return ctx.Status(400).SendString("<h2>Error: Date must be within the last 7 days</h2>")
		}
	}

	activities, err := h.activityService.GetActivitiesForDay(userID, date)
	if err != nil {
		return ctx.Status(500).SendString(fmt.Sprintf("<h2>Error fetching activities: %v</h2>", err))
	}

	fmt.Printf("activities: %+v\n", activities)
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

	fmt.Printf("checkin: %+v\n", checkin)
	return render.Render(ctx, views.DailyCheckinGet(ToDailyReportForm(checkin)))
}

func ToDailyReportForm(serviceDailyReport service.ServiceDailyReport) service.DailyReportForm {
	ret := service.DailyReportForm{}

	if serviceDailyReport.Date != "" {
		ret.Date = serviceDailyReport.Date[:10]
	} else {
		ret.Date = "Entries not found for selected date"
	}

	if serviceDailyReport.DidMeditate != false {
		ret.DidMeditate = fmt.Sprintf("%t", serviceDailyReport.DidMeditate)
	}

	if serviceDailyReport.MinutesOfSports != 0 {
		ret.MinutesOfSports = fmt.Sprintf("%d", serviceDailyReport.MinutesOfSports)
	}

	if serviceDailyReport.MealsEaten != 0 {
		ret.MealsEaten = fmt.Sprintf("%d", serviceDailyReport.MealsEaten)
	}

	if serviceDailyReport.WaterDrankLiters != 0 {
		ret.WaterDrankLiters = fmt.Sprintf("%f", serviceDailyReport.WaterDrankLiters)
	}

	if serviceDailyReport.StepsMade != 0 {
		ret.StepsMade = fmt.Sprintf("%d", serviceDailyReport.StepsMade)
	}

	if serviceDailyReport.SleepScore != 0 {
		ret.SleepScore = fmt.Sprintf("%d", serviceDailyReport.SleepScore)
	}

	if serviceDailyReport.HappinessRating != 0 {
		ret.HappinessRating = fmt.Sprintf("%d", serviceDailyReport.HappinessRating)
	}

	if serviceDailyReport.ProductivityScore != 0 {
		ret.ProductivityScore = fmt.Sprintf("%d", serviceDailyReport.ProductivityScore)
	}

	if serviceDailyReport.StressLevel != 0 {
		ret.StressLevel = fmt.Sprintf("%d", serviceDailyReport.StressLevel)
	}
	if serviceDailyReport.SocialInteractions != 0 {
		ret.SocialInteractions = fmt.Sprintf("%f", serviceDailyReport.SocialInteractions)
	}

	if serviceDailyReport.ScreenTimeHours != 0 {
		ret.ScreenTimeHours = fmt.Sprintf("%f", serviceDailyReport.ScreenTimeHours)
	}

	if serviceDailyReport.WorkHours != 0 {
		ret.WorkHours = fmt.Sprintf("%f", serviceDailyReport.WorkHours)
	}

	if serviceDailyReport.LeisureTimeHours != 0 {
		ret.LeisureTimeHours = fmt.Sprintf("%f", serviceDailyReport.LeisureTimeHours)
	}

	if serviceDailyReport.AlcoholUnits != 0 {
		ret.AlcoholUnits = fmt.Sprintf("%f", serviceDailyReport.AlcoholUnits)
	}

	if serviceDailyReport.CaffeineCups != 0 {
		ret.CaffeineCups = fmt.Sprintf("%f", serviceDailyReport.CaffeineCups)
	}

	if serviceDailyReport.OutdoorTimeHours != 0 {
		ret.OutdoorTimeHours = fmt.Sprintf("%f", serviceDailyReport.OutdoorTimeHours)
	}

	return ret
}
