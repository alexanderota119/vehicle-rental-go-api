package vehicle

import (
	"fmt"
	"os"

	"github.com/rfauzi44/vehicle-rental/database/orm/model"
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

	data.Image = fmt.Sprintf("%s%s%s", os.Getenv("BASE_URL"), "/public/image", data.Image)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *vehicle_repo) GetAll() (*model.Vehicles, error) {
	var data model.Vehicles
	err := r.database.Find(&data).Error
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
	data.Image = fmt.Sprintf("%s%s%s", os.Getenv("BASE_URL"), "/public/image", data.Image)

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
