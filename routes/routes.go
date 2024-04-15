package routes

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"sync"
)

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

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

func SetRoutes(app *fiber.App) {
	app.Get("/", AuthMiddleware, DashboardHandler)
	app.Post("/", AuthMiddleware, DashboardPostHandler)
	app.Post("/logout", LogoutHandler)

	app.Get("/login", LoginHandler)
	app.Post("/login", LoginPostHandler)

	//app.Get("/create", func(ctx *fiber.Ctx) error {
	//	u := &db.User{}
	//	u.CreateAdmin()
	//	return ctx.SendString("Created user")
	//})

	app.Get("/combined", func(c *fiber.Ctx) error {
		globalTimeData.Lock()
		defer globalTimeData.Unlock()

		return render(c, views.CombinedView(globalTimeData.Data))
	})

	app.Get("/total_times", func(c *fiber.Ctx) error {
		globalTimeData.Lock()
		defer globalTimeData.Unlock()

		return render(c, views.TotalTimes(totalTimeData.Data))
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
