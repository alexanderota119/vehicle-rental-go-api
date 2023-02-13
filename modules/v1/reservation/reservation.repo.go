package reservation

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"gorm.io/gorm"
)

type reservation_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *reservation_repo {
	return &reservation_repo{db}

}

func (r *reservation_repo) Add(data *model.Reservation) (*model.Reservation, error) {
	var vehicle model.Vehicle

	err := r.database.Where("vehicle_id = ?", data.FVehicleID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	data.Total = data.Duration * data.Quantity * vehicle.Price

	err = r.database.Create(&data).Preload("VehicleDetail").Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *reservation_repo) GetAll() (*model.Reservations, error) {
	var data model.Reservations
	err := r.database.Preload("VehicleDetail").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *reservation_repo) Update(data *model.Reservation) (*model.Reservation, error) {
	var vehicle model.Vehicle

	err := r.database.Where("vehicle_id = ?", data.FVehicleID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	data.Total = data.Duration * data.Quantity * vehicle.Price
	err = r.database.Model(&model.Reservation{}).Where("reservation_id = ?", data.ReservationID).Updates(data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *reservation_repo) Delete(uuid string) (*model.Reservation, error) {
	var data model.Reservation
	result := r.database.Where("reservation_id = ?", uuid).Delete(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}
