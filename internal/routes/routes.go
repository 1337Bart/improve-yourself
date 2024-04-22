package routes

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/handlers/login"
	settingsHandler "github.com/1337Bart/improve-yourself/internal/handlers/settings"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"sync"
)

var globalTimeData = struct {
	sync.Mutex
	Data views.TimeData
}{
	Data: views.TimeData{
		TimePool: 30,
	},
}

var totalTimeData = struct {
	sync.Mutex
	Data views.TotalTimeData
}{

	Data: views.TotalTimeData{
		Productivity: 0,
		Potato:       0,
	},
}

type potatoTimeForm struct {
	PotatoTime int `form:"potato"`
}

type productivityTimeForm struct {
	ProductivityTime int `form:"productivity"`
}

func SetRoutes(app *fiber.App, loginHandler *login.Handler, settingsHandler *settingsHandler.Handler) {
	app.Get("/", loginHandler.Index)
	app.Get("/settings", login.AuthMiddleware, settingsHandler.SettingsGet)
	app.Post("/settings", login.AuthMiddleware, settingsHandler.SettingsPost)

	app.Post("/logout", loginHandler.Logout)
	app.Get("/login", loginHandler.Login)
	app.Post("/login", loginHandler.LoginPost)

	app.Post("/register-user", loginHandler.RegisterUser)
	app.Post("/register-admin", loginHandler.RegisterAdmin)

	app.Get("/combined", func(c *fiber.Ctx) error {
		globalTimeData.Lock()
		defer globalTimeData.Unlock()

		return render.Render(c, views.CombinedView(globalTimeData.Data))
	})

	app.Get("/total_times", func(c *fiber.Ctx) error {
		globalTimeData.Lock()
		defer globalTimeData.Unlock()

		return render.Render(c, views.TotalTimes(totalTimeData.Data))
	})

	app.Post("/potato-time", func(c *fiber.Ctx) error {
		input := potatoTimeForm{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}

		globalTimeData.Lock()
		globalTimeData.Data.TimePool -= input.PotatoTime
		totalTimeData.Data.Potato -= input.PotatoTime
		globalTimeData.Unlock()

		fmt.Println(input)
		return c.SendStatus(200)
	})

	app.Post("/productivity-time", func(c *fiber.Ctx) error {
		input := productivityTimeForm{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}

		globalTimeData.Lock()
		globalTimeData.Data.TimePool += input.ProductivityTime
		totalTimeData.Data.Productivity += input.ProductivityTime
		globalTimeData.Unlock()

		fmt.Println(input)
		return c.SendStatus(200)
	})
}
