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
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
