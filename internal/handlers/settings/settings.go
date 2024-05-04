package settings

import (
	"fmt"
	db "github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	settingsService service.Settings
}

func NewHandler(s service.Settings) *Handler {
	return &Handler{settingsService: s}
}

func (h Handler) SettingsGet(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}
	fmt.Println("userId", userID)
	settings, err := h.settingsService.Get(userID)
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve settings</h2>")
	}
	amount := strconv.FormatUint(uint64(settings.Amount), 10)
	return render.Render(ctx, views.Home(amount, settings.SearchOn, settings.AddNew))
}

type settingsForm struct {
	Amount   int    `form:"amount"`
	SearchOn string `form:"searchOn"`
	AddNew   string `form:"addNew"`
}

func (h Handler) SettingsPost(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized access</h2>")
	}

	input := settingsForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve settings</h2>")
	}

	addNew := input.AddNew == "on"
	searchOn := input.SearchOn == "on"

	settings := &db.Settings{
		UUID:     userID,
		Amount:   uint(input.Amount),
		SearchOn: searchOn,
		AddNew:   addNew,
	}

	err := h.settingsService.Update(settings)
	if err != nil {
		fmt.Println(err)
		return ctx.SendString("<h2>Error: cannot update settings</h2>")
	}

	ctx.Append("HX-Refresh", "true")

	return ctx.SendStatus(200)
}
