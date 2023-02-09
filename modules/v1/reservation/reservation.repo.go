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

	err := r.database.Where("vehicle_id = ?", data.VehicleID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	data.Total = data.Duration * data.Quantity * vehicle.Price

	err = r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *reservation_repo) GetAll() (*model.Reservation, error) {
	var data model.Reservation
	err := r.database.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *reservation_repo) Update(data *model.Reservation) (*model.Reservation, error) {
	var vehicle model.Vehicle

	err := r.database.Where("vehicle_id = ?", data.VehicleID).First(&vehicle).Error
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

func (r *reservation_repo) Delete(uuid string) error {
	err := r.database.Where("reservation_id = ?", uuid).Delete(&model.Reservation{}).Error

	if err != nil {
		return err
	}

	return nil

}
