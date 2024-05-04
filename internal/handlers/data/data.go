package data

import (
	"github.com/1337Bart/improve-yourself/internal/service"
)

type Handler struct {
	settingsService service.Settings
	dataService     service.Data
}

func NewHandler(s service.Settings, d service.Data) *Handler {
	return &Handler{
		settingsService: s,
		dataService:     d,
	}
}

type TimePoolForm struct {
	Time int `form:"time"`
}

//
//func (h Handler) Login(ctx *fiber.Ctx) error {
//	return render.Render(ctx, views.Login())
//}
//
//func (h Handler) LoginPost(ctx *fiber.Ctx) error {
//	input := LoginForm{}
//	if err := ctx.BodyParser(&input); err != nil {
//		ctx.Status(500)
//		return ctx.SendString("<h2>Error: Something went wrong</h2>")
//	}
//
//	user := &model.User{}
//	user, err := h.loginService.LoginAsUser(input.Email, input.Password, user)
//	if err != nil {
//		ctx.Status(401)
//		return ctx.SendString("<h2>Error: Unauthorized</h2>")
//	}
//
//	signedToken, err := hashing.CreateNewAuthToken(user.ID, user.Email, user.IsAdmin)
//	if err != nil {
//		ctx.Status(401)
//		return ctx.SendString("<h2>Error: Something went wrong logging in</h2>")
//	}
//
//	cookie := fiber.Cookie{
//		Name:     "user_token",
//		Value:    signedToken,
//		Expires:  time.Now().Add(24 * time.Hour),
//		HTTPOnly: true,
//	}
//
//	ctx.Cookie(&cookie)
//	ctx.Append("HX-Redirect", "/")
//
//	return ctx.SendStatus(200)
//}
