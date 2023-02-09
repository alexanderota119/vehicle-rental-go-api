package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type ReservationRepoIF interface {
	Add(data *model.Reservation) (*model.Reservation, error)
	GetAll() (*model.Reservation, error)
	Update(data *model.Reservation) (*model.Reservation, error)
	Delete(uuid string) error
}

type ReservationServiceIF interface {
	Add(data *model.Reservation) *lib.Response
	GetAll() *lib.Response
	Update(data *model.Reservation) *lib.Response
	Delete(uuid string) *lib.Response
}
