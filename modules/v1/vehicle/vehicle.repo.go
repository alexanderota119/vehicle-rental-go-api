package vehicle

import (
	"fmt"

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

	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return data, nil

}

func (r *vehicle_repo) GetAll() (*model.Vehicles, error) {
	var data model.Vehicles
	err := r.database.Order("name desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil

}

func (r *vehicle_repo) GetById(uuid string) (*model.Vehicle, error) {
	var data model.Vehicle

	err := r.database.First(&data, "vehicle_id = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return &data, nil
}

func (r *vehicle_repo) Update(data *model.Vehicle) (*model.Vehicle, error) {
	err := r.database.Model(&model.Vehicle{}).Where("vehicle_id = ?", data.VehicleID).Updates(data).Error

	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

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

func (r *vehicle_repo) GetByCategory(category string) (*model.Vehicles, error) {
	var data model.Vehicles

	fmt.Println(category)

	err := r.database.Where("category = ?", category).Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil
}
