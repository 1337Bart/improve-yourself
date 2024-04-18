package login

import (
	"errors"
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	SqlDb *gorm.DB
}

func NewLoginService(sqlDbConn *gorm.DB) *Login {
	return &Login{
		SqlDb: sqlDbConn,
	}
}

func (l *Login) CreateAdmin() error {
	user := model.User{
		Email:    "your email",
		Password: "your password",
		IsAdmin:  true,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %s", err)
	}

	user.Password = string(password)

	err = l.SqlDb.Create(&user).Error
	if err != nil {
		return fmt.Errorf("CreateAdmin error: %s", err)
	}

	return nil
}

func (l *Login) LoginAsAdmin(email, password string, u *model.User) (*model.User, error) {
	if err := l.SqlDb.Where("email = ? AND is_admin = ?", email, true).First(&u).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return u, nil
}
