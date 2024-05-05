package data

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	settingsService service.Settings
	dataService     service.Data
}

func NewHandler(s service.Settings, d service.Data) *Handler {
	return &Handler{
		settingsService: s,
		dataService:     d,
	}
}

type TimePoolForm struct {
	Time   uint   `form:"time"`
	Action string `form:"action"`
}

func (h Handler) PotatoTimeGet(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	time, err := h.dataService.GetPotatoTime(userID)
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve potato time</h2>")
	}

	timeStr := strconv.Itoa(time)

	return render.Render(ctx, views.PotatoTime(timeStr))
}

func (h Handler) PotatoTimePost(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	input := TimePoolForm{}
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendString("<h2>Error: Unable to capture input</h2>")
	}

	var err error
	switch input.Action {
	case "add":
		err = h.dataService.AddPotatoTime(userID, input.Time)
	case "subtract":
		err = h.dataService.SubtractPotatoTime(userID, input.Time)
	default:
		return ctx.SendString("<h2>Error: Invalid input</h2>")
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("<h2>Error processing your request</h2>")
	}

	updatedTime, err := h.dataService.GetPotatoTime(userID)
	if err != nil {
		return ctx.Status(500).SendString("<h2>Error: Cannot retrieve updated time</h2>")
	}
	updatedTimeStr := strconv.Itoa(updatedTime)

	return ctx.SendString(fmt.Sprintf("<span id='productivity-time-counter'>%s</span>", updatedTimeStr))
}

func (h Handler) Dashboard(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	timesUpdated, err := h.dataService.GetPotatoTimeUpdatesCount(userID)
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve timesUpdated time</h2>")
	}

	totalUsedPotatoTime, err := h.dataService.GetTotalUsedTime(userID)
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve totalUsedPotatoTime time</h2>")
	}

	totalAddedProductivityTime, err := h.dataService.GetTotalAddedTime(userID)
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve totalAddedProductivityTime time</h2>")
	}

	return render.Render(ctx, views.Dashboard(int(totalAddedProductivityTime), int(totalUsedPotatoTime), int(timesUpdated)))
}
