package settings

import (
	"fmt"
	db "github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/routes"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func SettingsHandler(ctx *fiber.Ctx) error {
	settings := db.Settings{}
	err := settings.Get()
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve settings</h2>")
	}
	amount := strconv.FormatUint(uint64(settings.Amount), 10)
	return routes.Render(ctx, views.Home(amount, settings.SearchOn, settings.AddNew))
}

type settingsForm struct {
	Amount   int    `form:"amount"`
	SearchOn string `form:"searchOn"`
	AddNew   string `form:"addNew"`
}

func DashboardPostHandler(ctx *fiber.Ctx) error {
	input := settingsForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve settings</h2>")
	}

	// TODO - optionals here?
	addNew := false
	if input.AddNew == "on" {
		addNew = true
	}

	searchOn := false
	if input.SearchOn == "on" {
		searchOn = true
	}

	settings := &db.Settings{}
	settings.Amount = uint(input.Amount)
	settings.SearchOn = searchOn
	settings.AddNew = addNew

	err := settings.Update()
	if err != nil {
		fmt.Println(err)
		return ctx.SendString("<h2>Error: cannot update settings</h2>")
	}

	ctx.Append("HX-Refresh", "true")

	return ctx.SendStatus(200)
}
