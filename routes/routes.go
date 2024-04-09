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

type loginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type settingsForm struct {
	Amount   int  `form:"amount"`
	SearchOn bool `form:"searchOn"`
	AddNew   bool `form:"addNew"`
}

type potatoTimeForm struct {
	PotatoTime int `form:"potato"`
}

type productivityTimeForm struct {
	ProductivityTime int `form:"productivity"`
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})

	app.Post("/", func(c *fiber.Ctx) error {
		input := settingsForm{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}
		return c.SendStatus(200)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return render(c, views.Login())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		input := loginForm{}
		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}
		return c.SendStatus(200)
	})

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
