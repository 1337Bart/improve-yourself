package service

type Data interface {
	CreateNilPotatoTime(id string) error

	GetPotatoTime(id string) (time int, err error)

	AddPotatoTime(id string, amount uint) error
	SubtractPotatoTime(id string, amount uint) error

	GetPotatoTimeUpdatesCount(id string) (time uint, err error)
	GetTotalAddedTime(id string) (time uint, err error)
	GetTotalUsedTime(id string) (time uint, err error)
}
