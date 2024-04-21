package login

import (
	"github.com/1337Bart/improve-yourself/hashing"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/routes"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Handler struct {
	loginService service.Login
}

func NewHandler(l service.Login) *Handler {
	return &Handler{loginService: l}
}

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	return routes.Render(ctx, views.Login())
}

func (h Handler) LoginPost(ctx *fiber.Ctx) error {
	input := LoginForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: Something went wrong</h2>")
	}

	// tu chcę uzyc login, servis bedzie wybieral czy dostaje admina czy regular
	user := &model.User{}
	user, err := h.loginService.LoginAsAdmin(input.Email, input.Password, user)
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

func (h Handler) Logout(ctx *fiber.Ctx) error {
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
