package activity

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h *Handler) ActivitiesForDayGet(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		return ctx.Status(401).SendString("<h2>Error: Unauthorized access</h2>")
	}

	date := ctx.Query("selected_date")
	if date == "" {
		date = time.Now().Format("2006-01-02") // Format date to match HTML input element
	}

	activities, err := h.activityService.GetActivitiesForDay(userID, date)
	if err != nil {
		return ctx.Status(500).SendString(fmt.Sprintf("<h2>Error fetching activities: %v</h2>", err))
	}

	// Render the activities view, ensuring that this render function or method is correctly set up to handle the activities
	return render.Render(ctx, views.ActivityDayLog(activities))
}
