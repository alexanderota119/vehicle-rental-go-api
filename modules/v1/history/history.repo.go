package history

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"gorm.io/gorm"
)

type history_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}

}

func (r *history_repo) GetById(uuid string) (*model.Reservations, error) {
	var data model.Reservations
	err := r.database.Preload("VehicleDetail").Where("f_user_id = ?", uuid).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}
