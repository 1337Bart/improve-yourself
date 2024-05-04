package routes

import (
	dataHandler "github.com/1337Bart/improve-yourself/internal/handlers/data"
	"github.com/1337Bart/improve-yourself/internal/handlers/login"
	settingsHandler "github.com/1337Bart/improve-yourself/internal/handlers/settings"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, loginHandler *login.Handler, settingsHandler *settingsHandler.Handler, dataHandler *dataHandler.Handler) {
	app.Get("/", login.AuthMiddleware, loginHandler.Index)

	app.Post("/logout", loginHandler.Logout)
	app.Get("/login", loginHandler.Login)
	app.Post("/login", loginHandler.LoginPost)

	app.Get("/register-user", loginHandler.RegisterUser)
	app.Post("/register-user", loginHandler.RegisterUserPost)

	app.Get("/register-admin", loginHandler.RegisterAdmin)
	app.Post("/register-admin", loginHandler.RegisterAdminPost)

	app.Get("/settings", login.AuthMiddleware, settingsHandler.SettingsGet)
	app.Post("/settings", login.AuthMiddleware, settingsHandler.SettingsPost)

	app.Get("/potato-time", login.AuthMiddleware, dataHandler.PotatoTimeGet)
	app.Post("/update-time", login.AuthMiddleware, dataHandler.PotatoTimePost)

}

//app.Get("/total_times", func(c *fiber.Ctx) error {
//	globalTimeData.Lock()
//	defer globalTimeData.Unlock()
//
//	return render.Render(c, views.TotalTimes(totalTimeData.Data))
//})
//
//app.Post("/potato-time", func(c *fiber.Ctx) error {
//	input := potatoTimeForm{}
//	if err := c.BodyParser(&input); err != nil {
//		return c.SendString("<h2>Error: Something went wrong</h2>")
//	}
//
//	globalTimeData.Lock()
//	globalTimeData.Data.TimePool -= input.PotatoTime
//	totalTimeData.Data.Potato -= input.PotatoTime
//	globalTimeData.Unlock()
//
//	fmt.Println(input)
//	return c.SendStatus(200)
//})
