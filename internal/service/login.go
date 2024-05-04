package service

import "github.com/1337Bart/improve-yourself/internal/db/model"

type Login interface {
	CreateAdmin(string, string) error
	LoginAsAdmin(string, string, *model.User) (*model.User, error)

	CreateUser(string, string) error
	LoginAsUser(string, string, *model.User) (*model.User, error)

	GetUUIDByEmail(string) (string, error)
}
