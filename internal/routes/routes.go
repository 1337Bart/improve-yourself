package routes

import (
	activityHandler "github.com/1337Bart/improve-yourself/internal/handlers/activity"
	dataHandler "github.com/1337Bart/improve-yourself/internal/handlers/dashboard"
	"github.com/1337Bart/improve-yourself/internal/handlers/login"
	settingsHandler "github.com/1337Bart/improve-yourself/internal/handlers/settings"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, loginHandler *login.Handler, settingsHandler *settingsHandler.Handler, dataHandler *dataHandler.Handler, activityHandler *activityHandler.Handler) {
	app.Get("/", login.AuthMiddleware, dataHandler.Dashboard)

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

	app.Get("/dashboard", login.AuthMiddleware, dataHandler.Dashboard)

	app.Get("/activity-log", login.AuthMiddleware, activityHandler.LogActivityGet)
	app.Post("/activity-log", login.AuthMiddleware, activityHandler.LogActivityPost)

	app.Get("/activities-for-day", login.AuthMiddleware, activityHandler.ActivitiesForDayGet)
	app.Post("/activities-for-day", login.AuthMiddleware, activityHandler.ActivitiesForDayPost)

	app.Get("/daily-checkin", login.AuthMiddleware, activityHandler.LogDailyReportGet)
	app.Post("/daily-checkin", login.AuthMiddleware, activityHandler.LogDailyReportPost)

	app.Get("/daily-checkin-for-day", login.AuthMiddleware, activityHandler.DailyCheckinForDayGet)
}
