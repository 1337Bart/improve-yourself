package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ActivityLogDisplay struct {
	Activity  string
	StartTime string
	EndTime   string
	Duration  string
	Comments  string
}

// musze sie nauczyc przekazywac zakres dat
// bedzie oddawalo wszystkie aktywnosci dla zakresu dat (max 7 dni walidacja - inaczej sie rozjade)
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

	return render.Render(ctx, views.ActivityDayLog(activities))
}
