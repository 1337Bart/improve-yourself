package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/1337Bart/improve-yourself/db"
	"github.com/1337Bart/improve-yourself/hashing"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type loginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func LoginHandler(ctx *fiber.Ctx) error {
	return render(ctx, views.Login())
}

func LoginPostHandler(ctx *fiber.Ctx) error {
	input := loginForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: Something went wrong</h2>")
	}

	user := &db.User{}
	user, err := user.LoginAsAdmin(input.Email, input.Password)
	if err != nil {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Unauthorized</h2>")
	}

	signedToken, err := hashing.CreateNewAuthToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		ctx.Status(401)
		return ctx.SendString("<h2>Error: Something went wrong logging in</h2>")
	}

	cookie := fiber.Cookie{
		Name:     "admin",
		Value:    signedToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	ctx.Append("HX-Redirect", "/")

	return ctx.SendStatus(200)
}

func LogoutHandler(ctx *fiber.Ctx) error {
	ctx.ClearCookie("admin")
	ctx.Set("HX-Redirect", "/login")
	return ctx.SendStatus(200)
}

type AdminClaims struct {
	User                 string `json:"user"`
	Id                   string `json:"id"`
	jwt.RegisteredClaims `json:"claims"`
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("admin")
	if cookie == "" {
		return ctx.Redirect("/login", 302)
	}

	token, err := jwt.ParseWithClaims(cookie, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return ctx.Redirect("/login", 302)
	}

	_, ok := token.Claims.(*AdminClaims)
	if ok && token.Valid {
		return ctx.Next()
	}

	return ctx.Redirect("/login", 302)
}

func DashboardHandler(ctx *fiber.Ctx) error {
	settings := db.Settings{}
	err := settings.Get()
	if err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: cannot retrieve settings</h2>")
	}
	amount := strconv.FormatUint(uint64(settings.Amount), 10)
	return render(ctx, views.Home(amount, settings.SearchOn, settings.AddNew))
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
