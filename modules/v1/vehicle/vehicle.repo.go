package vehicle

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"gorm.io/gorm"
)

type vehicle_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}

}

func (r *vehicle_repo) GetAll() (*model.Vehicles, error) {
	var data model.Vehicles
	err := r.database.Model(model.Vehicle{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *vehicle_repo) Add(data *model.Vehicle) (*model.Vehicle, error) {
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
