package service

import "github.com/1337Bart/improve-yourself/internal/db/model"

type Settings interface {
	Get(string) (*model.Settings, error)
	Update(*model.Settings) error
}
