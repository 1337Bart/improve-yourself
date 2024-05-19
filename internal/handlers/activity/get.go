package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"time"
)

// musze sie nauczyc przekazywac zakres dat
// bedzie oddawalo wszystkie aktywnosci dla zakresu dat (max 7 dni walidacja - inaczej sie rozjade)
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
	} else {
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
