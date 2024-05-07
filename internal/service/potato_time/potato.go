package potato_time

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"gorm.io/gorm"
	"time"
)

type Data struct {
	SqlDb *gorm.DB
}

func NewDataService(sqlDbConn *gorm.DB) *Data {
	return &Data{
		SqlDb: sqlDbConn,
	}
}

type PotatoTime struct {
	ID           uint      `gorm:"autoIncrement;column:id" json:"id"`
	UUID         string    `gorm:"primaryKey;type:uuid;uniqueIndex;column:uuid" json:"uuid"`
	Amount       int       `json:"amount"`
	UpdatesCount uint      `json:"updatesCount"`
	TotalAdded   uint      `json:"totalAdded"`
	TotalUsed    uint      `json:"totalUsed"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (d *Data) AddPotatoTime(id string, amount uint) error {
	potatoTime := &model.PotatoTime{}
	err := d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error
	if err != nil {
		return fmt.Errorf("error fetching record: %s", err)
	}

	potatoTime.Amount += int(amount)
	potatoTime.UpdatesCount++
	potatoTime.TotalAdded += amount

	tx := d.SqlDb.Model(&model.PotatoTime{}).Where("uuid = ?", id).Updates(map[string]interface{}{
		"amount":        potatoTime.Amount,
		"updates_count": potatoTime.UpdatesCount,
		"total_added":   potatoTime.TotalAdded,
	})
	if err := tx.Error; err != nil {
		return fmt.Errorf("error updating record: %v", err)
	}

	return err
}

func (d *Data) SubtractPotatoTime(id string, amount uint) error {
	potatoTime := &model.PotatoTime{}
	err := d.SqlDb.Where("uuid = ?", id).First(&potatoTime).Error

	potatoTime.Amount -= int(amount)
	potatoTime.UpdatesCount++
	potatoTime.TotalUsed += amount

	tx := d.SqlDb.Model(&model.PotatoTime{}).Where("uuid = ?", id).Updates(map[string]interface{}{
		"amount":        potatoTime.Amount,
		"updates_count": potatoTime.UpdatesCount,
		"total_used":    potatoTime.TotalUsed,
	})
	if err := tx.Error; err != nil {
		return fmt.Errorf("error updating record: %v", err)
	}

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
