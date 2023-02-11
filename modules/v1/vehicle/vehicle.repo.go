package vehicle

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
	"gorm.io/gorm"
)

type vehicle_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}

}

func (r *vehicle_repo) Add(data *model.Vehicle) (*model.Vehicle, error) {
	err := r.database.Create(data).Error

	data.Image = lib.ImageReturn(data.Image)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *vehicle_repo) GetAll() (*model.Vehicles, error) {
	var data model.Vehicles
	err := r.database.Find(&data).Error

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *vehicle_repo) GetById(uuid string) (*model.Vehicle, error) {
	var data model.Vehicle

	err := r.database.First(&data, "vehicle_id = ?", uuid).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *vehicle_repo) Update(data *model.Vehicle) (*model.Vehicle, error) {
	err := r.database.Model(&model.Vehicle{}).Where("vehicle_id = ?", data.VehicleID).Updates(data).Error
	data.Image = lib.ImageReturn(data.Image)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *vehicle_repo) Delete(uuid string) (*model.Vehicle, error) {

	var data model.Vehicle
	result := r.database.Where("vehicle_id = ?", uuid).Delete(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}
