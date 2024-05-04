package data

import (
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"gorm.io/gorm"
)

type Data struct {
	SqlDb *gorm.DB
}

func NewDataService(sqlDbConn *gorm.DB) *Data {
	return &Data{
		SqlDb: sqlDbConn,
	}
}

func (d *Data) AddPotatoTime(id string, amount uint) error {
	potatoTime := &model.PotatoTime{}
	err := d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	currentPotatoTime := potatoTime.Amount
	updatesCount := potatoTime.UpdatesCount + 1
	totalAddedTime := potatoTime.TotalAdded + amount

	potatoTime.Amount = currentPotatoTime + int(amount)
	potatoTime.UpdatesCount = updatesCount
	potatoTime.TotalAdded = totalAddedTime

	tx := d.SqlDb.Select("amount", "updateCounts").Where("uuid = ?", id).Updates(potatoTime)
	err = tx.Error

	return err
}

func (d *Data) SubtractPotatoTime(id string, amount uint) error {
	potatoTime := &model.PotatoTime{}
	err := d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	currentPotatoTime := potatoTime.Amount
	updatesCount := potatoTime.UpdatesCount + 1
	totalUsedTime := potatoTime.TotalUsed + amount

	potatoTime.Amount = currentPotatoTime - int(amount)
	potatoTime.UpdatesCount = updatesCount
	potatoTime.TotalUsed = totalUsedTime

	tx := d.SqlDb.Select("amount", "updateCounts").Where("uuid = ?", id).Updates(potatoTime)
	err = tx.Error

	return err
}

func (d *Data) GetPotatoTime(id string) (time int, err error) {
	potatoTime := &model.PotatoTime{}
	err = d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	return potatoTime.Amount, err
}

func (d *Data) GetPotatoTimeUpdatesCount(id string) (time uint, err error) {
	potatoTime := &model.PotatoTime{}
	err = d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	return potatoTime.UpdatesCount, err
}

func (d *Data) CreateNilPotatoTime(id string) error {
	potatoTime := &model.PotatoTime{
		UUID:         id,
		Amount:       0,
		UpdatesCount: 0,
	}
	tx := d.SqlDb.Create(potatoTime)

	return tx.Error
}

func (d *Data) GetTotalUsedTime(id string) (time uint, err error) {
	potatoTime := &model.PotatoTime{}
	err = d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	return potatoTime.TotalUsed, err
}

func (d *Data) GetTotalAddedTime(id string) (time uint, err error) {
	potatoTime := &model.PotatoTime{}
	err = d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	return potatoTime.TotalAdded, err
}
