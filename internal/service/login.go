package service

import "github.com/1337Bart/improve-yourself/internal/db/model"

type Login interface {
	CreateAdmin() error
	LoginAsAdmin(string, string, *model.User) (*model.User, error)

	CreateUser(string, string) error
	LoginAsUser(string, string, *model.User) (*model.User, error)
}
