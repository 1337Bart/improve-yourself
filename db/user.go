package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	IsAdmin   bool      `gorm:"default:false" json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) CreateAdmin() error {
	user := User{
		Email:    "your email",
		Password: "your password",
		IsAdmin:  true,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %s", err)
	}

	user.Password = string(password)

	err = DbConn.Create(&user).Error
	if err != nil {
		return fmt.Errorf("CreateAdmin error: %s", err)
	}

	return nil
}

func (u *User) LoginAsAdmin(email, password string) (*User, error) {
	if err := DbConn.Where("email = ? AND is_admin = ?", email, true).First(&u).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return u, nil
}
