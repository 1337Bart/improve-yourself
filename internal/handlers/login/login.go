package login

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/hashing"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/render"
	"github.com/1337Bart/improve-yourself/internal/service"
	"github.com/1337Bart/improve-yourself/views"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Handler struct {
	loginService    service.Login
	settingsService service.Settings
	dataService     service.PotatoTime
}

func NewHandler(l service.Login, s service.Settings, d service.PotatoTime) *Handler {
	return &Handler{
		loginService:    l,
		settingsService: s,
		dataService:     d,
	}
}

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	return render.Render(ctx, views.Login())
}

func (h Handler) LoginPost(ctx *fiber.Ctx) error {
	input := LoginForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(500)
		return ctx.SendString("<h2>Error: Something went wrong</h2>")
	}

	user := &model.User{}
	user, err := h.loginService.LoginAsUser(input.Email, input.Password, user)
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
		Name:     "user_token",
		Value:    signedToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	ctx.Append("HX-Redirect", "/potato-time")

	return ctx.SendStatus(200)
}

func (h Handler) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("user_token")
	ctx.Set("HX-Redirect", "/login")
	return ctx.SendStatus(200)
}

type UserClaims struct {
	Email                string `json:"email"`
	ID                   string `json:"ID"`
	jwt.RegisteredClaims `json:"claims"`
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("user_token")
	if cookie == "" {
		return ctx.Redirect("/login", 302)
	}

	token, err := jwt.ParseWithClaims(cookie, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return ctx.Redirect("/login", 302)
	}

	claims, ok := token.Claims.(*UserClaims)
	if ok && token.Valid {
		ctx.Locals("userID", claims.ID)
		return ctx.Next()
	}

	return ctx.Redirect("/login", 302)
}

func (h Handler) RegisterAdmin(ctx *fiber.Ctx) error {
	return render.Render(ctx, views.RegisterAdmin())
}

func (h Handler) RegisterAdminPost(ctx *fiber.Ctx) error {
	input := LoginForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString("Error parsing input")
	}

	err := h.loginService.CreateAdmin(input.Email, input.Password)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to create admin: %s", err))
	}

	err = h.settingsService.CreateDefault(input.Email)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to initiate settings for user: %s", err))
	}

	err = h.dataService.CreateNilPotatoTime(input.Email)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to initiate potato time for user: %s", err))
	}

	return ctx.SendString("Admin created successfully")
}

func (h Handler) RegisterUser(ctx *fiber.Ctx) error {
	return render.Render(ctx, views.RegisterUser())
}

func (h Handler) RegisterUserPost(ctx *fiber.Ctx) error {
	input := LoginForm{}
	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString("Error parsing input")
	}

	err := h.loginService.CreateUser(input.Email, input.Password)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to create user: %s", err))
	}

	uuid, err := h.loginService.GetUUIDByEmail(input.Email)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to fetch uuid for user: %s", err))
	}

	err = h.settingsService.CreateDefault(uuid)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to initiate settings for user: %s", err))
	}

	err = h.dataService.CreateNilPotatoTime(uuid)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(fmt.Sprintf("Failed to initiate potato time for user: %s", err))
	}

	return ctx.SendString("User created successfully")
}
